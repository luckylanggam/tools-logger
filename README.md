# Logger

**Logger** is a package that prints messages of information, warnings, errors, etc. We use this package as a resource to obtain logs of things that are happening and that information helps us debug and know where in the code an error is occurring. This package is totally independent and the information it prints depends on what we pass to it.

![Go](https://img.shields.io/badge/Golang-1.21-blue.svg?logo=go&longCache=true&style=flat)

This package makes use of the external package [zap](go.uber.org/zap).

## Getting Started

[Go](https://golang.org/) is required in version 1.21 or higher.

### Install

`go get -u github.com/luckylanggam/tools-logger`

### Features

* [x] **Lightweight**, less than 200 lines of code.
* [x] **Easy** to use.

### Basic use

```go
package main

import (
	"github.com/luckylanggam/tools-logger"
	"os"
)

func main() {
	l := logger.New("service name", conf.Debug)

	//reading a file example
	filepath := "/filepath"
	
	l.Info("reading file at:", filepath)
	dat, err := os.ReadFile(filepath)
	if err != nil{
		l.With("file_error", "reading file").Error(err)
		os.Exit(1)
    }
	processFile(l, dat)

}

func processFile(l logger.Logger, _ []byte){
	l.Info("processing file")
	//process data
}

```
