//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: main.go
//     description:
//         created: 2015-12-19 10:03:52
//          author: wystan
//
//===============================================================================

//=======================================
// 1. export GO15VENDOREXPERIMENT=1
// 2. go build
// or
// 1. env export GO15VENDOREXPERIMENT=1 go build
//=======================================
package main

import (
	"fmt"

	"src/github.com/wy/wystan"
)

func main() {
	wystan.RunWystan()
	fmt.Printf("hello\n")
}

//==================================== END ======================================
