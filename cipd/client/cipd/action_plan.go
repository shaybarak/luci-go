// Copyright 2018 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cipd

import (
	"sort"

	"golang.org/x/net/context"

	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/logging"

	"go.chromium.org/luci/cipd/common"
)

// Helper structs and implementation guts of EnsurePackages call.

// Actions is returned by EnsurePackages.
//
// It lists pins that were attempted to be installed, updated or removed, as
// well as all errors.
type Actions struct {
	ToInstall common.PinSlice `json:"to_install,omitempty"` // pins to be installed
	ToUpdate  []UpdatedPin    `json:"to_update,omitempty"`  // pins to be replaced
	ToRemove  common.PinSlice `json:"to_remove,omitempty"`  // pins to be removed
	Errors    []ActionError   `json:"errors,omitempty"`     // all individual errors
}

// Empty is true if there are no actions specified.
func (a *Actions) Empty() bool {
	return len(a.ToInstall) == 0 && len(a.ToUpdate) == 0 && len(a.ToRemove) == 0
}

// UpdatedPin specifies a pair of pins: old and new version of a package.
type UpdatedPin struct {
	From common.Pin `json:"from"`
	To   common.Pin `json:"to"`
}

// ActionError holds an error that happened when installing or removing the pin.
type ActionError struct {
	Action string     `json:"action"`
	Pin    common.Pin `json:"pin"`
	Error  JSONError  `json:"error,omitempty"`
}

// ActionMap is a map of subdir to the Actions which will occur within it.
type ActionMap map[string]*Actions

// LoopOrdered loops over the ActionMap in sorted order (by subdir).
func (am ActionMap) LoopOrdered(cb func(subdir string, actions *Actions)) {
	subdirs := make(sort.StringSlice, 0, len(am))
	for subdir := range am {
		subdirs = append(subdirs, subdir)
	}
	subdirs.Sort()
	for _, subdir := range subdirs {
		cb(subdir, am[subdir])
	}
}

// Log prints the pending action to the logger installed in ctx.
func (am ActionMap) Log(ctx context.Context) {
	keys := make([]string, 0, len(am))
	for key := range am {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, subdir := range keys {
		actions := am[subdir]

		if subdir == "" {
			logging.Infof(ctx, "In root:")
		} else {
			logging.Infof(ctx, "In subdir %q:", subdir)
		}

		if len(actions.ToInstall) != 0 {
			logging.Infof(ctx, "  to install:")
			for _, pin := range actions.ToInstall {
				logging.Infof(ctx, "    %s", pin)
			}
		}
		if len(actions.ToUpdate) != 0 {
			logging.Infof(ctx, "  to update:")
			for _, pair := range actions.ToUpdate {
				logging.Infof(ctx, "    %s (%s -> %s)",
					pair.From.PackageName, pair.From.InstanceID, pair.To.InstanceID)
			}
		}
		if len(actions.ToRemove) != 0 {
			logging.Infof(ctx, "  to remove:")
			for _, pin := range actions.ToRemove {
				logging.Infof(ctx, "    %s", pin)
			}
		}
	}
}

// buildActionPlan is used by EnsurePackages to figure out what to install or
// remove.
func buildActionPlan(desired, existing common.PinSliceBySubdir) (aMap ActionMap) {
	desiredSubdirs := stringset.New(len(desired))
	for desiredSubdir := range desired {
		desiredSubdirs.Add(desiredSubdir)
	}

	existingSubdirs := stringset.New(len(existing))
	for existingSubdir := range existing {
		existingSubdirs.Add(existingSubdir)
	}

	aMap = ActionMap{}

	// all newly added subdirs
	desiredSubdirs.Difference(existingSubdirs).Iter(func(subdir string) bool {
		if want := desired[subdir]; len(want) > 0 {
			aMap[subdir] = &Actions{ToInstall: want}
		}
		return true
	})

	// all removed subdirs
	existingSubdirs.Difference(desiredSubdirs).Iter(func(subdir string) bool {
		if have := existing[subdir]; len(have) > 0 {
			aMap[subdir] = &Actions{ToRemove: have}
		}
		return true
	})

	// all common subdirs
	desiredSubdirs.Intersect(existingSubdirs).Iter(func(subdir string) bool {
		a := Actions{}

		// Figure out what needs to be installed or updated.
		haveMap := existing[subdir].ToMap()
		for _, want := range desired[subdir] {
			if haveID, exists := haveMap[want.PackageName]; !exists {
				a.ToInstall = append(a.ToInstall, want)
			} else if haveID != want.InstanceID {
				a.ToUpdate = append(a.ToUpdate, UpdatedPin{
					From: common.Pin{PackageName: want.PackageName, InstanceID: haveID},
					To:   want,
				})
			}
		}

		// Figure out what needs to be removed.
		wantMap := desired[subdir].ToMap()
		for _, have := range existing[subdir] {
			if wantMap[have.PackageName] == "" {
				a.ToRemove = append(a.ToRemove, have)
			}
		}

		if !a.Empty() {
			aMap[subdir] = &a
		}
		return true
	})

	if len(aMap) == 0 {
		return nil
	}
	return
}