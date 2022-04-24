# mntinfo

[![PkgGoDev](https://pkg.go.dev/badge/github.com/thediveo/go-mntinfo)](https://pkg.go.dev/github.com/thediveo/go-mntinfo)
[![GitHub](https://img.shields.io/github/license/thediveo/go-mntinfo)](https://img.shields.io/github/license/thediveo/go-mntinfo)
![build and test](https://github.com/TheDiveO/go-mntinfo/workflows/build%20and%20test/badge.svg?branch=master)
![file descriptors](https://img.shields.io/badge/file%20descriptors-not%20leaking-success)
[![Go Report Card](https://goreportcard.com/badge/github.com/TheDiveO/go-mntinfo)](https://goreportcard.com/report/github.com/TheDiveO/go-mntinfo)

`mntinfo` is a _minimalistic_ Linux-only Go package for discovering the
currently mounted filesystems seen by processes. This package also supports
discovering only those mounts matching a specific filesystem type.

> **Note:** mount discovery is done using `/proc/[PID]/mountinfo` data from
> Linux' `proc` filesystem.

## Copyright and License

`mntinfo` is Copyright 2019 Harald Albrecht, and licensed under the [Apache
License, Version 2.0](LICENSE).

## Requirements

Linux.

For a multi-platform solution please take a look at
[gopsutil](https://github.com/shirou/gopsutil) instead.
