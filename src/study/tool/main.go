//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: main.go
//     description:
//         created: 2015-12-16 09:18:50
//          author: wystan
//
//===============================================================================

// there are 3 init files here: run_linux.go run_windows.go run_darwin.go
// but only the init() in run_darwin.go will be linked into the final elf file
// that is because `go too dist` will ignore other os specific files;
// please refer to the file $GO_SRC/cmd/dist/build.go

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello\n")
}

//==================================== END ======================================
