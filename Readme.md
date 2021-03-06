# Human readable bytes 
[![Build Status](https://travis-ci.org/gangleri/humanbytes.svg?branch=master)](https://travis-ci.org/gangleri/humanbytes)
[![Report card](https://goreportcard.com/badge/gangleri.io/pkg/humanbytes)](https://goreportcard.com/report/gangleri.io/pkg/humanbytes)
[![codecov](https://codecov.io/gh/gangleri/humanbytes/branch/master/graph/badge.svg)](https://codecov.io/gh/gangleri/humanbytes)
[![GoDoc](https://godoc.org/gangleri.io/pkg/humanbytes?status.svg)](http://godoc.org/gangleri.io/pkg/humanbytes)

This package is used to convert bytes to and from human readable representations. It can convert a string such as
'1 MB' into bytes. Convert bytes to another unit eg KB, MB, GB... and also generate a human readable representation of
the bytes e.g. 1024 would be displayed as '1 KB'.

## Installation
```sh
go get -u gangleri.io/pkg/humanbytes

```

## Documentation

### Parsing
To convert a string such as `1 KB` into bytes:

```go
b, err := humanbytes.ParseBytes("1 KB") // 1024
```

### Convert
To convert the bytes into another unit e.g. `KB`:

```go
kb, err := humanbytes.Convert(1048576, "MB") // 1MB
```

### Sprint

```go
s, err := humanbytes.Sprint(1024, "KB") // 1KB
```

## Licence
BSD-2-Clause


