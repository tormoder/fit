# fit

[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/tormoder/fit/raw/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/tormoder/fit?status.svg)](https://godoc.org/github.com/tormoder/fit)
[![Travis Build Status](https://travis-ci.org/tormoder/fit.svg?branch=master)](https://travis-ci.org/tormoder/fit)

<img src="https://raw.githubusercontent.com/hackraft/gophericons/master/png/2.png" width="225" align="right" hspace="25" />

**This library is at the moment not actively maintained.**

fit is a [Go](http://www.golang.org/) package that implements decoding of the
[Flexible and Interoperable Data Transfer (FIT)
Protocol](http://www.thisisant.com/resources/fit). Fit is a "compact binary
format designed for storing and sharing data from sport, fitness and health
devices". Fit files are created by newer GPS enabled Garmin sport watches and
cycling computers, such as the Forerunner/Edge/Fenix series.

The two latest versions of Go is supported. The core decoding package has no
external dependencies. The latest release of Go and a few external dependencies
are required for running the full test suite and benchmarks.

### FIT SDK Version Support

**Current supported FIT SDK version:** 20.88

**Warning:** Data Developers Fields are not yet supported (#21).

Older supported profile versions:

* 20.43
* 20.27
* 20.14
* 16.20

Other profile versions may work, but have not been tested.

### Features

* Supports all FIT file types.
* Accessors for scaled fields.
* Accessors for dynamic fields.
* Field components expansion.
* Go code generation for custom FIT product profiles.

### Installation

```
$ go get github.com/tormoder/fit
```

### About fit

- [Example Usage](https://github.com/tormoder/fit/wiki/Example-Usage)
- [Data Types](https://github.com/tormoder/fit/wiki/Data-Types)
- [Main API Reference](https://github.com/tormoder/fit/wiki/Main-Api-Reference)
- [Custom Product Profiles](https://github.com/tormoder/fit/wiki/Custom-Product-Profiles)
- [Upcoming Features](https://github.com/tormoder/fit/wiki/Upcoming-Features)
- [Contributing](https://github.com/tormoder/fit/blob/master/CONTRIBUTING.md)
- [License](https://github.com/tormoder/fit/wiki/License)
