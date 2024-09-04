# vela-git

[![license](https://img.shields.io/crates/l/gl.svg)](../LICENSE)
[![GoDoc](https://godoc.org/github.com/go-vela/vela-git?status.svg)](https://godoc.org/github.com/go-vela/vela-git)
[![Go Report Card](https://goreportcard.com/badge/go-vela/vela-git)](https://goreportcard.com/report/go-vela/vela-git)
[![codecov](https://codecov.io/gh/go-vela/vela-git/branch/main/graph/badge.svg)](https://codecov.io/gh/go-vela/vela-git)

A Vela plugin designed for cloning repositories into your workspace.

Internally, the plugin is a wrapper around the [git](https://git-scm.com/) CLI.

The plugin comes in two flavors: `vela-git-slim` and `vela-git`

`vela-git-slim` contains basic packages required to perform basic git operations - ideal for regular clone steps.

`vela-git` contains all of `vela-git-slim` and additional utility packages related to git, such as `git-lfs` and `gh` (GitHub CLI).

## Documentation

For installation and usage, please [visit our user docs](https://go-vela.github.io/docs).

## Contributing

We are always welcome to new pull requests!

Please see our [contributing](CONTRIBUTING.md) documentation for further instructions.

## Support

We are always here to help!

Please see our [support](SUPPORT.md) documentation for further instructions.

## Copyright and License

```
Copyright 2019 Target Brands, Inc.
```

[Apache License, Version 2.0](../LICENSE)
