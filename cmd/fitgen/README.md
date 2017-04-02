# fitgen

```fitgen``` is a program that given a FIT profile generates Go code for
[fit](https://github.com/tormoder/fit). It takes as input the official FIT profile
specification workbook and outputs Go type, message and profile definitions.

Consult the [Wiki](https://github.com/tormoder/fit/wiki/Profile-Generation)
for information about profile generation.

## Prerequisites

* ```$GOPATH``` set.
* FIT SDK Zip or workbook file available.

## Usage

```shell
usage: fitgen [flags] [path to sdk zip, xls or xlsx file]
  -install
        run go install before invoking stringer (go/types related, see golang issue #11415)
  -sdk string
        provide or override SDK version printed in generated code
  -test
        run all tests in fit repository after code has been generated
  -timestamp
        add generation timestamp to generated code (default true)
```
