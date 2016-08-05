// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package server

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"

	"github.com/luci/gae/filter/txnBuf"
	"github.com/luci/gae/service/datastore"
	"github.com/luci/gae/service/info"
	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/common/errors"
	"github.com/luci/luci-go/common/logging"
	"github.com/luci/luci-go/server/auth"
	"github.com/luci/luci-go/server/auth/identity"
)

// SessionStore stores auth sessions in the datastore (always in the default
// namespace). It implements auth.SessionStore.
type SessionStore struct {
	Prefix string // used as prefix for datastore keys
}

// defaultNS returns GAE context configured to use default namespace.
func defaultNS(c context.Context) context.Context {
	return info.Get(c).MustNamespace("")
}

// noTxnDS returns datastore interface configured to escape any current
// transaction.
func noTxnDS(c context.Context) datastore.Interface {
	return txnBuf.GetNoTxn(c)
}

// OpenSession create a new session for a user with given expiration time.
// It returns unique session ID.
func (s *SessionStore) OpenSession(c context.Context, userID string, u *auth.User, exp time.Time) (string, error) {
	if strings.IndexByte(s.Prefix, '/') != -1 {
		return "", fmt.Errorf("gaeauth: bad prefix (%q) in SessionStore", s.Prefix)
	}
	if strings.IndexByte(userID, '/') != -1 {
		return "", fmt.Errorf("gaeauth: bad userID (%q), cannot have '/' inside", userID)
	}
	if err := u.Identity.Validate(); err != nil {
		return "", fmt.Errorf("gaeauth: bad identity string (%q) - %s", u.Identity, err)
	}
	c = defaultNS(c)

	now := clock.Now(c).UTC()
	prof := profile{
		Identity:  string(u.Identity),
		Superuser: u.Superuser,
		Email:     u.Email,
		Name:      u.Name,
		Picture:   u.Picture,
	}

	// Set in the transaction below.
	var sessionID string

	err := noTxnDS(c).RunInTransaction(func(c context.Context) error {
		ds := datastore.Get(c)

		// Grab existing userEntity or initialize a new one.
		userEnt := userEntity{ID: s.Prefix + "/" + userID}
		err := ds.Get(&userEnt)
		if err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}
		if err == datastore.ErrNoSuchEntity {
			userEnt.Profile = prof
			userEnt.Created = now
		}
		userEnt.LastLogin = now

		// Make new session. ID will be generated by the datastore.
		sessionEnt := sessionEntity{
			Parent:     ds.KeyForObj(&userEnt),
			Profile:    prof,
			Created:    now,
			Expiration: exp.UTC(),
		}
		if err = ds.Put(&userEnt, &sessionEnt); err != nil {
			return err
		}

		sessionID = fmt.Sprintf("%s/%s/%d", s.Prefix, userID, sessionEnt.ID)
		return nil
	}, nil)

	if err != nil {
		return "", errors.WrapTransient(err)
	}
	return sessionID, nil
}

// CloseSession closes a session given its ID. Does nothing if session is
// already closed or doesn't exist. Returns only transient errors.
func (s *SessionStore) CloseSession(c context.Context, sessionID string) error {
	c = defaultNS(c)
	ent, err := s.fetchSession(c, sessionID)
	switch {
	case err != nil:
		return err
	case ent == nil:
		return nil
	default:
		ent.IsClosed = true
		ent.Closed = clock.Now(c).UTC()
		return errors.WrapTransient(noTxnDS(c).Put(ent))
	}
}

// GetSession returns existing non-expired session given its ID. Returns nil
// if session doesn't exist, closed or expired. Returns only transient errors.
func (s *SessionStore) GetSession(c context.Context, sessionID string) (*auth.Session, error) {
	c = defaultNS(c)
	ent, err := s.fetchSession(c, sessionID)
	if ent == nil {
		return nil, err
	}
	return &auth.Session{
		SessionID: sessionID,
		UserID:    ent.Parent.StringID()[len(s.Prefix)+1:],
		User: auth.User{
			Identity:  identity.Identity(ent.Profile.Identity),
			Superuser: ent.Profile.Superuser,
			Email:     ent.Profile.Email,
			Name:      ent.Profile.Name,
			Picture:   ent.Profile.Picture,
		},
		Exp: ent.Expiration,
	}, nil
}

// fetchSession fetches sessionEntity from the datastore and returns it if it is
// still open and non-expired. Returns (nil, nil) otherwise. Broken sessionID is
// logged and ignored, the function returns (nil, nil) in such case. Returns
// only transient errors.
func (s *SessionStore) fetchSession(c context.Context, sessionID string) (*sessionEntity, error) {
	chunks := strings.Split(sessionID, "/")
	if len(chunks) != 3 || chunks[0] != s.Prefix {
		logging.Warningf(c, "Malformed session ID %q, ignoring", sessionID)
		return nil, nil
	}
	id, err := strconv.ParseInt(chunks[2], 10, 64)
	if err != nil {
		logging.Warningf(c, "Malformed session ID %q, ignoring", sessionID)
		return nil, nil
	}
	ds := noTxnDS(c)
	sessionEnt := sessionEntity{
		ID:     id,
		Parent: ds.MakeKey("gaeauth.User", chunks[0]+"/"+chunks[1]),
	}
	switch err = ds.Get(&sessionEnt); err {
	case nil:
		if sessionEnt.IsClosed || clock.Now(c).After(sessionEnt.Expiration) {
			return nil, nil
		}
		return &sessionEnt, nil
	case datastore.ErrNoSuchEntity:
		return nil, nil
	default:
		return nil, errors.WrapTransient(err)
	}
}

////

// profile is used in both userEntity and sessionEntity. It holds information
// about a user extracted from user.User struct.
type profile struct {
	Identity  string
	Superuser bool   `gae:",noindex"`
	Email     string `gae:",noindex"`
	Name      string `gae:",noindex"`
	Picture   string `gae:",noindex"`
}

// userEntity holds profile information of some user. It is root entity.
// ID is "<prefix>/<userID>" where <prefix> is SessionStore.Prefix, and <userID>
// is what is passed to OpenSession (unique user id as returned by
// authentication backend). Created or refreshed in OpenSession.
type userEntity struct {
	_kind string `gae:"$kind,gaeauth.User"`

	ID string `gae:"$id"`

	Profile   profile
	Created   time.Time // when this entity was created
	LastLogin time.Time // when last session was opened
}

// sessionEntity stores session information associated with session cookie.
// Parent entity is userEntity, ID is generated by the datastore. Includes user
// profile info inline to avoid additional datastore calls in GetSession. Never
// deleted from the datastore (to keep some sort of history of logins). Marked
// as closed in CloseSession. Since all user's sessions belong to single entity
// group, there's implicit 1 login per second per user limit on rate of logins.
type sessionEntity struct {
	_kind string `gae:"$kind,gaeauth.Session"`

	ID     int64          `gae:"$id"`
	Parent *datastore.Key `gae:"$parent"`

	Profile    profile
	Created    time.Time // when this session was created
	Expiration time.Time // when this session expires

	IsClosed bool      // if true, the session was closed by CloseSession()
	Closed   time.Time // when the session was closed by CloseSession()
}
