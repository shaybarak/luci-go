// Copyright 2016 The LUCI Authors.
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

package common

import (
	"golang.org/x/net/context"

	"go.chromium.org/luci/luci_config/common/cfgtypes"
	"go.chromium.org/luci/luci_config/server/cfgclient/access"
	"go.chromium.org/luci/luci_config/server/cfgclient/backend"
	"go.chromium.org/luci/server/auth"
)

// Helper functions for ACL checking.

// IsAllowed checks to see if the user in the context is allowed to access
// the given project.
func IsAllowed(c context.Context, project string) (bool, error) {
	switch admin, err := IsAdmin(c); {
	case err != nil:
		return false, err
	case admin:
		return true, nil
	}
	// Get the project, because that's where the ACLs lie.
	err := access.Check(
		c, backend.AsUser,
		cfgtypes.ProjectConfigSet(cfgtypes.ProjectName(project)))
	switch err {
	case nil:
		return true, nil
	case access.ErrNoAccess:
		return false, nil
	default:
		return false, err
	}
}

// IsAdmin returns true if the current identity is an administrator.
func IsAdmin(c context.Context) (bool, error) {
	// TODO(nodir): unhardcode group name to config file if there is a need
	return auth.IsMember(c, "administrators")
}
