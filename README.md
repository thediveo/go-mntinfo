# mntinfo

[![mntinfo package
documentation](https://godoc.org/github.com/TheDiveO/go-mntinfo?status.svg)](http://godoc.org/github.com/TheDiveO/go-mntinfo)

`mntinfo` is a minimalistic Linux-only Go package for discovering the
currently mounted filesystems seen by processes. This package also supports
discovering only these mounts matching a specific filesystem type.

Discovery is done using `/proc/[PID]/mountinfo` data from Linux' `proc`
filesystem.

## Copyright and License

`mntinfo` is Copyright 2019 Harald Albrecht, and licensed under the Apache
License, Version 2.0.

## Requirements

Linux.

For a multi-platform solution please take a look at
[gopsutil](https://github.com/shirou/gopsutil) instead.
