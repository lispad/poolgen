## poolgen
Tool `poolgen` provides easy way to generate pool of objects, using [sync/pool](https://golang.org/src/sync/pool.go).

## install
```sh
$ go get -u github.com/lispad/poolgen
```

## usage
```sh
$ poolgen /path/to/the/go/file
```

Or use annotations in your golang code 
 
```golang
//go:generate poolgen $GOFILE
type MyStruct struct {
	Field1  string
	Field2  uint32
	Nested  SubStruct
	Pointer *SubStruct
}

type SubStruct struct {
	Field string
}

```

## Disclaimer
It's just a test project at the very beginning stage. Please, don't use for prod code.

This program is distributed WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE

