//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: main.go
//     description:
//         created: 2015-12-06 13:36:41
//          author: wystan
//
//===============================================================================

package main

import "fmt"

type TT struct {
	A int
}

func (t *TT) Run(c int) error {
	fmt.Printf("run %d\n", t.A+c)
	return nil
}

func (t *TT) Call(b int) error {
	fmt.Printf("call %d\n", t.A+b)
	return nil
}

type MethodCall func(b int) error

var myMap = make(map[int32]MethodCall)

func main() {
	fmt.Printf("hello\n")
	t := new(TT)
	t.A = 10
	myMap[1] = t.Run
	myMap[2] = t.Call

	myMap[1](10)
	myMap[2](4)
}

//==================================== END ======================================
