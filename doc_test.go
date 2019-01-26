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

// Lists the types of filesystems currently mounted, in alphabetical order.
// Most code in this example is needed for creating the final, sorted list of
// unique filesystem types.
func ExampleMounts_uniqueFstypes() {
	mounts := mntinfo.Mounts(-1)

	fstypes := map[string]bool{} // ...poor gopher's set emulation
	fstype_names := []string{}   // gather all unique fstypes
	for _, mount := range mounts {
		if _, ok := fstypes[mount.FsType]; !ok {
			fstypes[mount.FsType] = true
			fstype_names = append(fstype_names, mount.FsType)
		}
	}
	sort.Strings(fstype_names)
	fmt.Printf("currently mounted filesystem types: %s", strings.Join(fstype_names, ", "))
	// currently mounted filesystem types: autofs, binfmt_misc, cgroup, cgroup2, configfs, debugfs, ... vfat
}

// Lists only (bind-mounted) Linux namespaces. In Linux, namespaces are
// represented in the filesystem using the "nsfs" filesystem type. Tools like
// Docker or iproute2 (ip netns ...) bind-mount (network) namespaces in
// additional places outside the /proc filesystem in order to keep them open
// until the bind-mount is removed again, regardless of whether there are
// still any processes left using them.
func ExampleFsTypeMounts() {
	mounts := mntinfo.FsTypeMounts(-1, "nsfs")
	for _, mount := range mounts {
		fmt.Printf("namespace %s at %s", mount.Root, mount.MountPoint)
	}
	// namespace net:[4026532757] at /run/docker/netns/4281a40c1612
}
