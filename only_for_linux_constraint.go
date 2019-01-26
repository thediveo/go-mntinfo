// The only purpose of this go file is to cause a compilation error when
// trying to deploy this package on a non-Linux platform. Only on Linux, the
// symbol referenced will be defined by "only_for_linux.go". This technique
// has been inspired by: https://github.com/theckman/goconstraint

package mntinfo

var _ = package_mntinfo_requires_Linux
