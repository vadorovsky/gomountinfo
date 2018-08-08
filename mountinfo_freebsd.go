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

/*
#include <sys/param.h>
#include <sys/ucred.h>
#include <sys/mount.h>
*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Parse /proc/self/mountinfo because comparing Dev and ino does not work from
// bind mounts.
func ParseMountTable(filter FilterFunc) ([]*MountInfo, error) {
	var rawEntries *C.struct_statfs

	count := int(C.getmntinfo(&rawEntries, C.MNT_WAIT))
	if count == 0 {
		return nil, fmt.Errorf("Failed to call getmntinfo")
	}

	var entries []C.struct_statfs
	header := (*reflect.SliceHeader)(unsafe.Pointer(&entries))
	header.Cap = count
	header.Len = count
	header.Data = uintptr(unsafe.Pointer(rawEntries))

	var out []*MountInfo
	for _, entry := range entries {
		var mountinfo MountInfo
		var skip, stop bool
		mountinfo.MountPoint = C.GoString(&entry.f_mntonname[0])

		if filter != nil {
			// filter out entries we're not interested in
			skip, stop = filter(p)

			if skip {
				continue
			}
		}

		mountinfo.MountSource = C.GoString(&entry.f_mntfromname[0])
		mountinfo.FilesystemType = C.GoString(&entry.f_fstypename[0])

		out = append(out, &mountinfo)
		if stop {
			break
		}
	}
	return out, nil
}
