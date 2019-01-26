// +build linux

// Only on Linux, this go file defined the required constraint variable
// allowing this package to be built successfully.

package mntinfo

// This constant ensures that this package can only be successfully build
// on/for the Linux platform. With all other platforms, this canary constant
// will be left out, so that trying to build "only_for_linux.go" will fail
// with some hopefully slightly useful error message.
const requiresLinux = 42
