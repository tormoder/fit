# fitgen

```fitgen``` is a program that given a FIT profile generates Go code for
[gofit](https://github.com/tormoder/gofit). It takes as input the official FIT profile
specification workbook and outputs Go type, message and profile definitions.

Consult the [Wiki](https://github.com/tormoder/gofit/wiki/Profile-Generation)
for information about profile generation.

```fitgen``` has only been tested on Linux x86-64.

## Prerequisites

* ```$GOPATH``` set.
* Python 3 available in ```$PATH``` (xlsx to csv conversion).
* [xlrd](https://pypi.python.org/pypi/xlrd) Python library (xlsx to csv conversion).  
* FIT SDK Zip file available.

## Usage

```shell
usage: fitgenprofile [-keep] [path to sdk zip, xls or xlsx file]
  -jmptable
	use jump tables for profile message and field lookups, otherwise use switches (default true)
  -keep
	don't delete intermediary workbook and csv files from profile directory
  -sdk string
	provide or override SDK version printed in generated code
```
