# transition
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hslam/transition)](https://pkg.go.dev/github.com/hslam/transition)
[![Build Status](https://github.com/hslam/transition/workflows/build/badge.svg)](https://github.com/hslam/transition/actions)
[![codecov](https://codecov.io/gh/hslam/transition/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/transition)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/transition)](https://goreportcard.com/report/github.com/hslam/transition)
[![LICENSE](https://img.shields.io/github/license/hslam/transition.svg?style=flat-square)](https://github.com/hslam/transition/blob/master/LICENSE)

Package transition implements smooth transition.

## Get started

### Install
```
go get github.com/hslam/transition
```
### Import
```
import "github.com/hslam/transition"
```
### Usage
#### Example
```go
package main

import (
	"fmt"
	"github.com/hslam/transition"
)

func main() {
	concurrency := func() int {
		return 32
	}
	trans := transition.NewTransition(4, concurrency)
	low := func() {}
	high := func() {
		fmt.Println("Hello World")
	}
	trans.Smooth(low, high)
}
```

#### Output
```
Hello World
```

### License
This package is licensed under a MIT license (Copyright (c) 2021 Meng Huang)

### Author
transition was written by Meng Huang.


