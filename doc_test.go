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

package mntinfo_test

import (
	"fmt"
	"sort"
	"strings"

	mntinfo "github.com/thediveo/go-mntinfo"
)

// Shows all mounted filesystems with lots of details.
func Example() {
	mounts := mntinfo.Mounts()
	for _, mount := range mounts {
		fmt.Printf("%v\n", mount)
	}
	// Output should be slightly similar to:
	//   {22 28 0 21 / /sys [rw nosuid nodev noexec relatime] map[shared:7] sysfs sysfs rw}
	//   {23 28 0 4 / /proc [rw nosuid nodev noexec relatime] map[shared:14] proc proc rw}
	//   {24 28 0 6 / /dev [rw nosuid relatime] map[shared:2] devtmpfs udev rw,size=8158448k,nr_inodes=2039612,mode=755}
	//   {25 24 0 22 / /dev/pts [rw nosuid noexec relatime] map[shared:3] devpts devpts rw,gid=5,mode=620,ptmxmode=000}
	//   {26 28 0 23 / /run [rw nosuid noexec relatime] map[shared:5] tmpfs tmpfs rw,size=1637744k,mode=755}
	//   {28 0 8 2 / / [rw relatime] map[shared:1] ext4 /dev/sda2 rw,errors=remount-ro,data=ordered}
	//   {29 22 0 7 / /sys/kernel/security [rw nosuid nodev noexec relatime] map[shared:8] securityfs securityfs rw}
	//   ...
}

// Lists all mounted filesystems: where they are mounted and of which fs type
// they are. This list is sorted by mount path.
func ExampleMounts_sortedPaths() {
	mounts := mntinfo.Mounts()
	sort.Slice(mounts, func(a, b int) bool {
		return mounts[a].MountPoint < mounts[b].MountPoint
	})
	for _, mount := range mounts {
		fmt.Printf("%s of type %s\n", mount.MountPoint, mount.FsType)
	}
}

// Lists the types of filesystems currently mounted, in alphabetical order.
// Most code in this example is needed for creating the final, sorted list of
// unique filesystem types.
func ExampleMounts_uniqueFstypes() {
	mounts := mntinfo.Mounts()

	fstypes := map[string]bool{} // ...poor gopher's set emulation
	fstypeNames := []string{}    // gather all unique fstypes
	for _, mount := range mounts {
		if _, ok := fstypes[mount.FsType]; !ok {
			fstypes[mount.FsType] = true
			fstypeNames = append(fstypeNames, mount.FsType)
		}
	}
	sort.Strings(fstypeNames)
	fmt.Printf("currently mounted filesystem types: %s", strings.Join(fstypeNames, ", "))
	// currently mounted filesystem types: autofs, binfmt_misc, cgroup, cgroup2,
	// configfs, debugfs, ... vfat
}

// Lists only (bind-mounted) Linux namespaces. In Linux, namespaces are
// represented in the filesystem using the "nsfs" filesystem type. Tools like
// Docker or iproute2 (ip netns ...) bind-mount (network) namespaces in
// additional places outside the /proc filesystem in order to keep them open
// until the bind-mount is removed again, regardless of whether there are
// still any processes left using them.
func ExampleMountsOfType() {
	mounts := mntinfo.MountsOfType(-1, "nsfs")
	for _, mount := range mounts {
		fmt.Printf("namespace %s at %s", mount.Root, mount.MountPoint)
	}
	// namespace net:[4026532757] at /run/docker/netns/4281a40c1612
}
