# mntinfo

[![PkgGoDev](https://pkg.go.dev/badge/github.com/thediveo/go-mntinfo)](https://pkg.go.dev/github.com/thediveo/go-mntinfo)
[![GitHub](https://img.shields.io/github/license/thediveo/go-mntinfo)](https://img.shields.io/github/license/thediveo/go-mntinfo)
![build and test](https://github.com/TheDiveO/go-mntinfo/workflows/build%20and%20test/badge.svg?branch=master)
![file descriptors](https://img.shields.io/badge/file%20descriptors-not%20leaking-success)
[![Go Report Card](https://goreportcard.com/badge/github.com/TheDiveO/go-mntinfo)](https://goreportcard.com/report/github.com/TheDiveO/go-mntinfo)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

`mntinfo` is a _minimalistic_ Linux-only Go package for discovering the
currently mounted filesystems seen by processes. This package additionally
supports discovering only those mounts matching a specific filesystem type.

> **Note:** mount discovery is done using `/proc/[PID]/mountinfo` data from the
> `proc` filesystem – see also
> [proc(5)](https://man7.org/linux/man-pages/man5/proc.5.html).

## Installation

```bash
go get github.com/thediveo/go-mntinfo
```

## Hacking It

- to view the package documentation locally:
  - either: `make pkgsite`, then navigate to http://localhost:6060/github.com/thediveo/go-plugger;
  - or, in VSCode (using the VSCode-integrated simple browser): “Tasks: Run
    Task” ⇢ “View Go module documentation”.
- `make` shows the available make targets.

## Copyright and License

`mntinfo` is Copyright 2019-23 Harald Albrecht, and licensed under the
[Apache License, Version 2.0](LICENSE).

## Requirements

Linux.

For a multi-platform solution please take a look at
[gopsutil](https://github.com/shirou/gopsutil) instead.
