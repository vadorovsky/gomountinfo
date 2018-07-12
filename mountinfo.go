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

// MountInfo is a struct representing information from /proc/pid/mountinfo in
// Linux. It has the following syntax:
//
// 36 35 98:0 /mnt1 /mnt2 rw,noatime master:1 - ext3 /dev/root rw,errors=continue
// (1)(2)(3)   (4)   (5)      (6)      (7)   (8) (9)   (10)         (11)
// (1) mount ID:  unique identifier of the mount (may be reused after umount)
// (2) parent ID:  ID of parent (or of self for the top of the mount tree)
// (3) major:minor:  value of st_dev for files on filesystem
// (4) root:  root of the mount within the filesystem
// (5) mount point:  mount point relative to the process's root
// (6) mount options:  per mount options
// (7) optional fields:  zero or more fields of the form "tag[:value]"
// (8) separator:  marks the end of the optional fields
// (9) filesystem type:  name of filesystem of the form "type[.subtype]"
// (10) mount source:  filesystem specific information or "none"
// (11) super options:  per super block options
//
// More information:
// https://www.kernel.org/doc/Documentation/filesystems/proc.txt
//
// The same structure is used for FreeBSD, because its mountinfo doesn't contain
// any other fields.
type MountInfo struct {
	MountID        int64
	ParentID       int64
	Major          int
	Minor          int
	Root           string
	MountPoint     string
	MountOptions   string
	OptionalFields []string
	FilesystemType string
	MountSource    string
	SuperOptions   []string
}
