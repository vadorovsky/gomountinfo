// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gomountinfo

import (
	"strings"
)

// FilterFunc is a type defining a callback function to filter out unwanted
// entries. It takes a pointer to an Info struct (not fully populated, currently
// only Mountpoint is filled in), and returns two booleans:
//  - skip: true if the entry should be skipped
//  - stop: true if parsing should be stopped after the entry
type FilterFunc func(*MountInfo) (skip, stop bool)

// PrefixFilter discards all entries whose mount points do not start with a
// prefix specified
func PrefixFilter(prefix string) FilterFunc {
	return func(m *MountInfo) (bool, bool) {
		skip := !strings.HasPrefix(m.MountPoint, prefix)
		return skip, false
	}
}

// SingleEntryFilter looks for a specific entry
func SingleEntryFilter(mp string) FilterFunc {
	return func(m *MountInfo) (bool, bool) {
		if m.MountPoint == mp {
			return false, true // don't skip, stop now
		}
		return true, false // skip, keep going
	}
}

// ParentsFilter returns all entries whose mount points can be parents of a path
// specified, discarding others. For example, given `/var/lib/docker/something`,
// entries like `/var/lib/docker`, `/var` and `/` are returned.
func ParentsFilter(path string) FilterFunc {
	return func(m *MountInfo) (bool, bool) {
		skip := !strings.HasPrefix(path, m.MountPoint)
		return skip, false
	}
}
