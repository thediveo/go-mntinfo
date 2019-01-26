// +build linux

// Only on Linux, this go file defined the required constraint variable
// allowing this package to be built successfully.

package mntinfo

// This constant ensures that this package can only be successfully build
// on/for the Linux platform. With all other platforms, this canary constant
// will be left out, so that trying to build "linux_only.go" will fail with
// some hopefully slightly useful error message.
const package_mntinfo_requires_Linux = 42
