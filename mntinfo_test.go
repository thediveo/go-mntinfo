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
	"encoding/json"
	"os"
	"testing/fstest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	. "github.com/thediveo/fdooze"
)

var _ = Describe("mntinfo", func() {

	When("parsing mountinfo lines", func() {

		DescribeTable("rejects malformed lines",
			func(mntline string) {
				Expect(parseProcMountinfoLine(mntline)).Error().To(HaveOccurred())
			},
			EntryDescription("%s"),
			Entry(nil, ""),
			Entry(nil, "abc"),
			Entry(nil, "42"),
			Entry(nil, "42 abc"),
			Entry(nil, "42 42"),
			Entry(nil, "42 42 foo"),
			Entry(nil, "42 42 foo:bar"),
			Entry(nil, "42 42 42:42"),
			Entry(nil, "42 42 abc:42"),
			Entry(nil, "42 42 42:abc"),
			Entry(nil, "42 42 42:42 /"),
			Entry(nil, "42 42 42:42 / /proc"),
			Entry(nil, "42 42 42:42 / /proc foo,bar,baz"),
			Entry(nil, "42 42 42:42 / /proc foo,bar,baz froz:42"),
			Entry(nil, "42 42 42:42 / /proc foo,bar,baz froz:42 baz:42"),
			Entry(nil, "42 42 42:42 / /proc foo,bar,baz froz:42 baz:42 -"),
			Entry(nil, "42 42 42:42 / /proc foo,bar,baz froz:42 baz:42 - abcfs"),
			Entry(nil, "42 42 42:42 / /proc foo,bar,baz froz:42 baz:42 - abcfs abcfs"),
			Entry(nil, "42 42 42:42 / /proc foo,bar,baz froz:42 baz:42 - abcfs abcfs "),
		)

		DescribeTable("returns correct mount information",
			func(mntline string, mntinfo Mountinfo) {
				Expect(parseProcMountinfoLine(mntline)).To(Equal(mntinfo))
			},
			EntryDescription("%s"),
			Entry(nil, "1 2 3:4 / /proc foo,bar - abcfs defs rw", Mountinfo{
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
			}),
			Entry(nil, "1 2 3:4 / /proc foo,bar frotz barz:42 - abcfs defs rw", Mountinfo{
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
			}),
		)

	})

	When("parsing procfs mount information", func() {

		It("returns an error when the process is missing", func() {
			emptyfs := fstest.MapFS{}
			Expect(parseProcMountinfo(emptyfs, 42)).To(BeEmpty())
		})

		It("returns mount info for our process, skipping garbage", func() {
			selffs := fstest.MapFS{
				"self/mountinfo": &fstest.MapFile{
					Data: []byte(`24 31 0:22 / /sys rw,nosuid,nodev,noexec,relatime shared:7 - sysfs sysfs rw
Linus had a little penguin
25 31 0:23 / /proc rw,nosuid,nodev,noexec,relatime shared:12 - proc proc rw
`),
				},
			}
			Expect(parseProcMountinfo(selffs, -1)).To(HaveLen(2))
		})

		It("returns mount info for a particular process", func() {
			selffs := fstest.MapFS{
				"42/mountinfo": &fstest.MapFile{
					Data: []byte(`24 31 0:22 / /sys rw,nosuid,nodev,noexec,relatime shared:7 - sysfs sysfs rw
25 31 0:23 / /proc rw,nosuid,nodev,noexec,relatime shared:12 - proc proc rw
`),
				},
			}
			Expect(parseProcMountinfo(selffs, 42)).To(HaveLen(2))
		})

	})

	When("parsing /proc/self/mountinfo", func() {

		BeforeEach(func() {
			goodfds := Filedescriptors()
			DeferCleanup(func() {
				Expect(Filedescriptors()).NotTo(HaveLeakedFds(goodfds))
			})
		})

		It("reads self mountinfo", func() {
			// There needs to be at least one mount for "/" on "/" ... or
			// otherwise something is really rotten here.
			minfo := Mounts()
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
			Expect(MountsOfPid(mypid)).To(Equal(Mounts()))
		})

		It("doesn't read from non-existing PID", func() {
			Expect(len(MountsOfPid(int(^uint(0) >> 1)))).To(BeZero())
		})

	})

	It("translates to/from JSON", func() {
		minfo := Mounts()[0]
		b, err := json.Marshal(minfo)
		Expect(err).ToNot(HaveOccurred())
		var minfo2 Mountinfo
		Expect(json.Unmarshal(b, &minfo2)).ToNot(HaveOccurred())
		Expect(minfo2).To(Equal(minfo))
	})

})
