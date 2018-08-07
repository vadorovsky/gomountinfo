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

package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	flag "github.com/spf13/pflag"

	"github.com/mrostecki/gomountinfo"
)

const (
	titleMountID        = "Mount ID"
	titleParentID       = "Parent ID"
	titleMajor          = "Major"
	titleMinor          = "Minor"
	titleRoot           = "Root"
	titleMountPoint     = "Mount point"
	titleMountOptions   = "Mount options"
	titleOptionalFields = "Optional fields"
	titleFilesystemType = "Filesystem type"
	titleMountSource    = "Mount source"
	titleSuperOptions   = "Super options"
)

var (
	pid int
)

func init() {
	flag.IntVarP(&pid, "pid", "p", -1, "PID of the process to get mountinfo from")
	flag.Parse()
}

func main() {
	w := tabwriter.NewWriter(os.Stdout, 5, 0, 3, ' ', 0)

	fmt.Fprintf(
		w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n",
		titleMountID, titleParentID, titleMajor, titleMinor, titleRoot,
		titleMountPoint, titleMountOptions, titleOptionalFields,
		titleFilesystemType, titleMountSource, titleSuperOptions,
	)

	var mountInfos []*gomountinfo.MountInfo
	var err error
	if pid > 0 {
		mountInfos, err = gomountinfo.ParseMountTablePid(pid, nil)
	} else {
		mountInfos, err = gomountinfo.ParseMountTable(nil)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get mount info: %v", err)
		os.Exit(1)
	}

	for _, mountInfo := range mountInfos {
		fmt.Fprintf(
			w, "%d\t%d\t%d\t%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n",
			mountInfo.MountID, mountInfo.ParentID, mountInfo.Major,
			mountInfo.Minor, mountInfo.Root, mountInfo.MountPoint,
			mountInfo.MountOptions, mountInfo.OptionalFields,
			mountInfo.FilesystemType, mountInfo.MountSource,
			mountInfo.SuperOptions,
		)
	}

	w.Flush()
}
