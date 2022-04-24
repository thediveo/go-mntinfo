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

//go:build linux
// +build linux

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thediveo/go-mntinfo"
)

func dumpNetNs(cmd *cobra.Command, _ []string) error {
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
	return nil
}

// newRootCmd creates the root command with usage and version information, as
// well as the available CLI flags (including descriptions).
func newRootCmd() (rootCmd *cobra.Command) {
	rootCmd = &cobra.Command{
		Use:     "findnetns",
		Short:   "findnetns finds all bind-mounted Linux network namespaces",
		Version: "0.9.1",
		Args:    cobra.NoArgs,
		RunE:    dumpNetNs,
	}
	// no additional CLI flags.
	return
}

func main() {
	// This is cobra boilerplate documentation, except for the missing call to
	// fmt.Println(err) which in the original boilerplate is just plain wrong:
	// it renders the error message twice, see also:
	// https://github.com/spf13/cobra/issues/304
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
