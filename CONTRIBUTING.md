# Contributing

I'm grateful for any help, whether it's code, documentation, or spelling
corrections.

## Filing issues

Please report at least Go version, operating system and processor architecture
when filing an issue.

## Contributing code

Please follow standard Go conventions and make sure that invoking ```make``` in
the root directory reports no errors.

## Adding test FIT files

When contributing a change to the parsing of FIT files it is beneficial to
include a FIT file which causes the issue you have resolved. To ensure the
FIT file is executed during tests you must do the following:

Include the FIT file in an appropriate subdirectory of "testdata" such as
"misc".

Add an entry to decodeTestFiles in the file reader_files_test.go, ensuring
that the fingerprint field is set to "1".

Execute ```go test -update``` to calculate the fingerprint and update it. This
will also produce a gzipped golden test output file in the same directory as
your original file. This file can be inspected by using a tool like ```zcat```.

Execute ```make test``` and ensure all tests have passed.

Commit your newly added FIT file, the generated gzipped golden test output
file, and the reader_files_test.go file.
