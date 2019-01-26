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

// +build linux

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	mntinfo "github.com/thediveo/go-mntinfo"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// nolint: unused,deadcode,varcheck
var (
	app     = kingpin.New(filepath.Base(os.Args[0]), "Finds all bind-mounted Linux network namespaces.")
	version = app.Version("0.9.0")
	_       = app.HelpFlag.Short('h') // now that's hidden deep inside the code...
)

func main() {
	_ = kingpin.MustParse(app.Parse(os.Args[1:]))
	mounts := mntinfo.MountsOfType(-1, "nsfs")
	// Sort all mount namespaces by their namespace ID.
	sort.Slice(mounts, func(a, b int) bool {
		return mounts[a].Root < mounts[b].Root
	})
	for _, mount := range mounts {
		if strings.HasPrefix(mount.Root, "net:[") {
			fmt.Printf("%s at %s\n", mount.Root, mount.MountPoint)
		}
	}
}
