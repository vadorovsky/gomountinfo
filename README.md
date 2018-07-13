# gomountinfo

[![Build Status](https://travis-ci.org/mrostecki/gomountinfo.svg?branch=master)](https://travis-ci.org/mrostecki/gomountinfo)

Go library for parsing information from /proc/self/mountinfo

## Motivation

There are many projects written in Go which make use of mountinfo. They are
implementing mountinfo parsers on their own. The purpose of this library is to
provide an universal way of getting mount information for Go programs.

## Credits to the previous implementations

This library is **not** written from scratch. It's heavily based on previously
existing implementations of mountinfo parser in:

* [Cilium](https://github.com/cilium/cilium)
* [Kubernetes](https://github.com/kubernetes/kubernetes)
* [Moby](https://github.com/moby/moby)

All those projects are licensed under Apache License 2.0 and authors of their
mountinfo modules are credited
[here](https://github.com/mrostecki/gomountinfo/blob/master/AUTHORS).
