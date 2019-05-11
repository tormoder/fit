# fitgen

`fitgen` is a program that given a FIT profile generates Go code for
[fit](https://github.com/tormoder/fit). It takes as input the official FIT profile
specification workbook and outputs Go type, message and profile definitions.

Consult the [Wiki](https://github.com/tormoder/fit/wiki/Profile-Generation)
for information about profile generation.

## Prerequisites

* FIT SDK Zip or workbook file available.

## Usage

```shell
usage: fitgen [flags] [path to sdk zip, xls or xlsx file] [output directory]
  -sdk string
        provide or override SDK version printed in generated code
  -test
        run all tests in output directory after code has been generated
  -timestamp
        add generation timestamp to generated code
  -verbose
        print verbose debugging output for profile parsing and code generation
```

## Global and product profiles

The complete specification of types and messages for the FIT protocol is called
the "Global Profile". A "Product Profile" is a subset of the Global Profile,
only including a subset of type and message definitions. A custom product
profile is a convenient way of controlling struct sizes and memory usage.

The `fitgen` tool supports generating code for custom product profiles. This
can be done by editing the SDK profile workbook and invoking
[fitgen](https://github.com/tormoder/fit/tree/master/cmd/fitgen).  Every field
definition with **a value greater than zero** in the ```EXAMPLE``` column of
the ```Messages``` sheet will be included in the generated code.

The fit package will always use the official FIT product profile bundled with
the SDK. It would be natural to fork the fit repository if you want to use your
own custom profile.

See Section 3, "*Overview of the FIT File protocol*" [1], for more information
about FIT product profiles.

## References

[1] Flexible & Interoperable Data Transfer (FIT) Protocol.
    [http://www.thisisant.com/resources/fit/](http://www.thisisant.com/resources/fit/)
