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
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("mntinfo", func() {

	Describe("parses mountinfo lines", func() {

		It("rejecting malformed lines", func() {
			malformed := []string{
				"",
				"abc", "42",
				"42 abc", "42 42",
				"42 42 foo", "42 42 foo:bar", "42 42 42:42", "42 42 abc:42", "42 42 42:abc",
				"42 42 42:42 /",
				"42 42 42:42 / /proc",
				"42 42 42:42 / /proc foo,bar,baz",
				"42 42 42:42 / /proc foo,bar,baz froz:42",
				"42 42 42:42 / /proc foo,bar,baz froz:42 baz:42",
				"42 42 42:42 / /proc foo,bar,baz froz:42 baz:42 -",
				"42 42 42:42 / /proc foo,bar,baz froz:42 baz:42 - abcfs",
				"42 42 42:42 / /proc foo,bar,baz froz:42 baz:42 - abcfs abcfs",
				"42 42 42:42 / /proc foo,bar,baz froz:42 baz:42 - abcfs abcfs ",
			}

			for _, malle := range malformed {
				_, err := parseProcMountinfoLine(malle)
				Expect(err).To(HaveOccurred(), "for line %q", malle)
			}
		})

		It("returning correct mount information", func() {
			Expect(parseProcMountinfoLine("1 2 3:4 / /proc foo,bar - abcfs defs rw")).To(Equal(Mountinfo{
				MountID:      1,
				ParentID:     2,
				Major:        3,
				Minor:        4,
				Root:         "/",
				MountPoint:   "/proc",
				MountOptions: []string{"foo", "bar"},
				Tags:         map[string]string{},
				FsType:       "abcfs",
				Source:       "defs",
				SuperOptions: "rw",
			}))
			Expect(parseProcMountinfoLine("1 2 3:4 / /proc foo,bar frotz barz:42 - abcfs defs rw")).To(Equal(Mountinfo{
				MountID:      1,
				ParentID:     2,
				Major:        3,
				Minor:        4,
				Root:         "/",
				MountPoint:   "/proc",
				MountOptions: []string{"foo", "bar"},
				Tags: map[string]string{
					"frotz": "",
					"barz":  "42",
				},
				FsType:       "abcfs",
				Source:       "defs",
				SuperOptions: "rw",
			}))
		})

	})

	Describe("parses /proc/self/mountinfo", func() {

		It("reads self mountinfo", func() {
			// There needs to be at least one mount for "/" on "/" ... or
			// otherwise something is really rotten here.
			minfo := MountsOfPid(-1)
			Expect(len(minfo)).NotTo(BeZero())
			Expect(minfo).To(ContainElement(
				MatchFields(IgnoreExtras, Fields{
					"Root":       Equal("/"),
					"MountPoint": Equal("/"),
				})))
		})

		It("filters mountinfo by fs type", func() {
			minfo := MountsOfType(-1, "proc")
			Expect(len(minfo)).NotTo(BeZero())
		})

		It("reads mountinfo from PID", func() {
			mypid := os.Getpid()
			Expect(MountsOfPid(mypid)).To(Equal(MountsOfPid(-1)))
		})

		It("doesn't read from non-existing PID", func() {
			Expect(len(MountsOfPid(int(^uint(0) >> 1)))).To(BeZero())
		})

	})

})
