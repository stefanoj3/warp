## Warp

[![Build Status](https://travis-ci.org/stefanoj3/warp.svg?branch=master)](https://travis-ci.org/stefanoj3/warp)
[![codecov](https://codecov.io/gh/stefanoj3/warp/branch/master/graph/badge.svg)](https://codecov.io/gh/stefanoj3/warp)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/stefanoj3/warp/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/stefanoj3/warp/?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/stefanoj3/warp)](https://goreportcard.com/report/github.com/stefanoj3/warp)
[![GoDoc](https://godoc.org/github.com/stefanoj3/warp?status.svg)](https://godoc.org/github.com/stefanoj3/warp)

Warp is a small library that abstracts Linux/macOS ARP table (read only).

Usage example:
```go
scanner, err := warp.NewARPScanner()
if err != nil {
    panic(err)
}

entries, err := scanner.Scan()
if err != nil {
    panic(err)
}

for _, entry := range entries {
    fmt.Println(entry)
}
```

You can check the full documentation at [godoc](https://godoc.org/github.com/stefanoj3/warp) or 
you can fetch the repo and run `make doc` to run a local godoc.

There is also an example file in the repo that you can look at for more "advanced" examples.

#### What is missing
Windows support is missing, it should not be too hard to introduce,
if you want it feel free to open a PR.


#### Contribution

Fork and submit a PR, tests and linters run in CI and you can also run them locally.
Run `make help` to see all the available commands (and their descriptions).