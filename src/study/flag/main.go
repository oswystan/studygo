//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: main.go
//     description:
//         created: 2015-11-28 07:49:34
//          author: wystan
//
//===============================================================================

package main

import (
	"flag"
	"fmt"
	"os"
)

type CmdOps struct {
	/* data */
	IntVar int
	StrVar string
}

func main() {
	var opt CmdOps
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.IntVar(&opt.IntVar, "int", 0, "set a int value")
	f.StringVar(&opt.StrVar, "str", "", "set a string value")
	f.Parse(os.Args[1:])

	fmt.Printf("%v\n", opt)
}

//==================================== END ======================================
