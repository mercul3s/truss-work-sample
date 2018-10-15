## CSV Normalization Work Sample

## Build and Run this sample
### Option 1:
* download and run a [release](https://github.com/mercul3s/truss-work-sample/releases) appropriate
for your OS. Builds are available for Linux, Windows, and macOS. 
* after you've untarred or unzipped the file, you can run the binary from the
`./bin` directory:
```
$ ./bin/normalize samples/sample.csv # or substitute the path to your file here
finished normalizing input file, writing to normalized_sample.csv
```
* Once the file has been processed, you can view the results in the normalized
csv file.

### Option 2:
* Clone this repo, build, and run locally.
* You will need to have [go](https://golang.org/doc/install) installed on your
machine. This codebase uses standard libraries, and should be compatible with go
1.x versions.
* run `dep ensure` to install library dependencies.
* run `go build .` to compile a binary called `work_sample`.
* run the binary on the command line, providing it a path to the file you want
to normalize: 
```
$ ./truss-work-sample samples/sample.csv # or substitute the path to your file here
finished normalizing input file, writing to normalized_sample.csv

```
* note: `go build .` will name the binary after the current working directory,
so if you clone this repo into another directory name, your binary may be
different than the example above.
* Once the file has been processed, you can view the results in the normalized
csv file.
