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
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	mountInfoFilepath = "/proc/self/mountinfo"
)

func parseInfoFile(r io.Reader, filter FilterFunc) ([]*MountInfo, error) {
	var result []*MountInfo

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("failed to scan the line from infofile: %v", err)
		}

		mountInfoRaw := scanner.Text()

		// Optional fields (which are on the 7th position) are separated
		// from the rest of fields by "-" character. The number of
		// optional fields can be greater or equal to 1.
		mountInfoSeparated := strings.Split(mountInfoRaw, " - ")
		if len(mountInfoSeparated) != 2 {
			return nil, fmt.Errorf("invalid mountinfo entry which has more that one '-' separator: %s", mountInfoRaw)
		}

		// Extract fields from both sides of mountinfo
		mountInfoLeft := strings.Split(strings.TrimSpace(mountInfoSeparated[0]), " ")
		// After '-' separator there should be 3 fields. The last field
		// may contain spaces.
		mountInfoRight := strings.SplitN(strings.TrimSpace(mountInfoSeparated[1]), " ", 3)

		// Before '-' separator there should be 6 fields and unknown
		// number of optional fields
		lenMountInfoLeft := len(mountInfoLeft)
		if lenMountInfoLeft < 6 {
			return nil, fmt.Errorf("invalid mountinfo entry, got %d fields before optional fields, exected at least 6: %s", lenMountInfoLeft, mountInfoRaw)
		}

		mountID, err := strconv.ParseInt(mountInfoLeft[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse mount ID: %v", err)
		}

		parentID, err := strconv.ParseInt(mountInfoLeft[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse parent ID: %v", err)
		}

		majorMinor := strings.Split(mountInfoLeft[2], ":")
		if len(majorMinor) != 2 {
			return nil, fmt.Errorf("unexpected minor:major pair '%s', invalid mountinfo entry: %s", mountInfoLeft[2], mountInfoRaw)
		}

		major, err := strconv.Atoi(majorMinor[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse major: %v", err)
		}

		minor, err := strconv.Atoi(majorMinor[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse minor: %v", err)
		}

		// Extract optional fields, which start from 7th position
		var optionalFields []string
		for i := 6; i < len(mountInfoLeft); i++ {
			optionalFields = append(optionalFields, mountInfoLeft[i])
		}

		// Split super options
		superOptions := strings.Split(mountInfoRight[2], ",")

		mountInfo := &MountInfo{
			MountID:        mountID,
			ParentID:       parentID,
			Major:          major,
			Minor:          minor,
			Root:           mountInfoLeft[3],
			MountPoint:     mountInfoLeft[4],
			MountOptions:   mountInfoLeft[5],
			OptionalFields: optionalFields,
			FilesystemType: mountInfoRight[0],
			MountSource:    mountInfoRight[1],
			SuperOptions:   superOptions,
		}

		// Filter out entries we're not interested in
		var skip, stop bool
		if filter != nil {
			skip, stop = filter(mountInfo)
			if skip {
				continue
			}
		}

		result = append(result, mountInfo)

		if stop {
			break
		}
	}

	return result, nil
}

func ParseMountTable(filter FilterFunc) ([]*MountInfo, error) {
	f, err := os.Open(mountInfoFilepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open mount info file %s: %v", mountInfoFilepath, err)
	}
	defer f.Close()

	return parseInfoFile(f, filter)
}
