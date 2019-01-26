// Copyright 2019 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mntinfo

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Mountinfo stores per-mount information as discovered from the proc
// filesystem. For mountinfo details, please refer to the proc(5) man page,
// and there to /proc/[PID]/mountinfo in particular.
type Mountinfo struct {
	// unique ID for the mount, might be reused after umount(2).
	MountID int `json:"mountid"`
	// ID of the parent mount, or self for the root of a mount namespace's mount tree.
	ParentID int `json:"parentid"`
	// major ID for the st_dev for files on this filesystem.
	Major int `json:"major"`
	// minor ID for the st_dev for filed on this filesystem.
	Minor int `json:"minor"`
	// pathname of the directory in the filesystem which forms the root of this mount.
	Root string `json:"root"`
	// pathname of the mount point relative to root directory of the process.
	MountPoint string `json:"mountpoint"`
	// mount options specific to this mount.
	MountOptions []string `json:"mountoptions"`
	// optional fields "tag[:value]"; tags cannot be a single hyphen "-".
	Tags map[string]string `json:"tags"`
	// filesystem type in the form "type[.subtype]"
	FsType string `json:"fstype"`
	// filesystem-specific information or "none".
	Source string `json:"source"`
	// per-superblock options.
	SuperOptions string `json:"superoptions"`
}

// Mounts returns all mounts for the current process.
func Mounts() []Mountinfo {
	return parseProcMountinfo(-1)
}

// MountsOfPid returns all mounts for either the current process (when pid specified
// as -1), or for another process identified by its PID.
func MountsOfPid(pid int) []Mountinfo {
	return parseProcMountinfo(pid)
}

// MountsOfType returns only those mounts for the current or another process
// (-1 or specific PID) matching the given fstype. Some fstypes are "ext4",
// "proc", "sysfs", "vfat", and many more.
func MountsOfType(pid int, fstype string) []Mountinfo {
	mounts := parseProcMountinfo(pid)
	filtered := []Mountinfo{}
	for idx := range mounts {
		if mounts[idx].FsType == fstype {
			filtered = append(filtered, mounts[idx])
		}
	}
	return filtered
}

// Fetches the mount information for a specific process (by PID) from the
// Linux kernel's procfs and parses it into a slice of Mountinfo elements, one
// for each mount.
func parseProcMountinfo(pid int) (mi []Mountinfo) {
	mi = []Mountinfo{}
	var pidstr string

	if pid <= 0 {
		pidstr = "self"
	} else {
		pidstr = strconv.Itoa(pid)
	}
	mif, err := os.Open(filepath.Join("/proc", pidstr, "mountinfo"))
	if err != nil {
		return
	}
	defer mif.Close()

	// Read in all lines from /proc/.../mountinfo, silently skipping any
	// garbage line we might encounter on our way.
	mifscan := bufio.NewScanner(mif)
	for mifscan.Scan() {
		mntline := mifscan.Text()
		if info, err := parseProcMountinfoLine(mntline); err == nil {
			mi = append(mi, info)
		}
	}
	return
}

// Parses a single line from /proc/[PID]/mountinfo, returning the information
// in a Mountinfo struct.
func parseProcMountinfoLine(mntline string) (info Mountinfo, err error) {
	// (1) mount ID
	info.MountID, mntline, err = nextInt(mntline)
	if err != nil {
		return
	}

	// (2) parent ID
	info.ParentID, mntline, err = nextInt(mntline)
	if err != nil {
		return
	}

	// (3) major:minor
	majmins, mntline, err := nextString(mntline)
	if err != nil {
		return
	}
	majmin := strings.Split(majmins, ":")
	if len(majmin) != 2 {
		err = errors.New("malformed major:minor field")
		return
	}
	major, err := strconv.Atoi(majmin[0])
	if err != nil {
		return
	}
	info.Major = major
	minor, err := strconv.Atoi(majmin[1])
	if err != nil {
		return
	}
	info.Minor = minor

	// (4) root
	info.Root, mntline, err = nextString(mntline)
	if err != nil {
		return
	}

	// (5) mount point
	info.MountPoint, mntline, err = nextString(mntline)
	if err != nil {
		return
	}

	// (6) mount options
	opts, mntline, err := nextString(mntline)
	if err != nil {
		return
	}
	info.MountOptions = strings.Split(opts, ",")

	// (7-8) optional fields, until single hyphen separator
	tags := map[string]string{}
	for {
		var tag string
		tag, mntline, err = nextString(mntline)
		if err != nil {
			return
		}
		if tag == "-" {
			break
		}
		namevalue := strings.SplitN(tag, ":", 2)
		if len(namevalue) < 2 {
			tags[namevalue[0]] = ""
		} else {
			tags[namevalue[0]] = namevalue[1]
		}
	}
	info.Tags = tags

	// (9) filesystem type
	info.FsType, mntline, err = nextString(mntline)
	if err != nil {
		return
	}

	// (10) mount source
	info.Source, mntline, err = nextString(mntline)
	if err != nil {
		return
	}

	// (11) super options
	info.SuperOptions, _, err = nextString(mntline)

	return
}

// Snipps off the first element from a string of space-delimited elements and
// returns it as an integer value, together with the remaining line for
// further processing.
func nextInt(line string) (i int, remline string, err error) {
	elems := strings.SplitN(line, " ", 2)
	if len(elems) >= 1 && len(elems[0]) > 0 {
		i, err = strconv.Atoi(elems[0])
		if len(elems) >= 2 {
			remline = elems[1]
		}
		return
	}
	return 0, "", errors.New("not enough elements in mountinfo line")
}

// Snipps off the first element from a string of space-delimited elements and
// returns it as a string value, together with the remaining line for further
// processing.
func nextString(line string) (s string, remline string, err error) {
	elems := strings.SplitN(line, " ", 2)
	if len(elems) >= 1 && len(elems[0]) > 0 {
		s = elems[0]
		if len(elems) >= 2 {
			remline = elems[1]
		}
		return
	}
	return "", "", errors.New("not enough elements in mountinfo line")
}
