# mntinfo

[![PkgGoDev](https://pkg.go.dev/badge/github.com/thediveo/go-mntinfo)](https://pkg.go.dev/github.com/thediveo/go-mntinfo)
[![GitHub](https://img.shields.io/github/license/thediveo/go-mntinfo)](https://img.shields.io/github/license/thediveo/go-mntinfo)
![build and test](https://github.com/thediveo/go-mntinfo/actions/workflows/buildandtest.yaml/badge.svg?branch=master)
![file descriptors](https://img.shields.io/badge/file%20descriptors-not%20leaking-success)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/go-mntinfo)](https://goreportcard.com/report/github.com/thediveo/go-mntinfo)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

`mntinfo` is a _minimalistic_ Linux-only Go package for discovering the
currently mounted filesystems seen by processes. This package additionally
supports discovering only those mounts matching a specific filesystem type.

> **Note:** mount discovery is done using `/proc/[PID]/mountinfo` data from the
> `proc` filesystem – see also
> [proc(5)](https://man7.org/linux/man-pages/man5/proc.5.html).

## Usage

```bash
go get github.com/thediveo/go-mntinfo
```

## DevContainer

> [!CAUTION]
>
> Do **not** use VSCode's "~~Dev Containers: Clone Repository in Container
> Volume~~" command, as it is utterly broken by design, ignoring
> `.devcontainer/devcontainer.json`.

1. `git clone https://github.com/thediveo/go-mntinfo`
2. in VSCode: Ctrl+Shift+P, "Dev Containers: Open Workspace in Container..."
3. select `go-mntinfo.code-workspace` and off you go...

## Supported Go Versions

`netdb` supports versions of Go that are noted by the [Go release
policy](https://golang.org/doc/devel/release.html#policy), that is, major
versions _N_ and _N_-1 (where _N_ is the current major version).

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md).

## Copyright and License

`mntinfo` is Copyright 2019-26 Harald Albrecht, and licensed under the
[Apache License, Version 2.0](LICENSE).

## Requirements

Linux.

For a multi-platform solution please take a look at
[gopsutil](https://github.com/shirou/gopsutil) instead.
