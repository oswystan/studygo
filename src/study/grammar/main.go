//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: main.go
//     description:
//         created: 2015-12-08 16:14:48
//          author: wystan
//
//===============================================================================

package main

import (
	"fmt"
	"time"
)

func breakLoop() {
	var i = 0
Loop:
	for i = 0; i < 10; i++ {
		switch {
		case i > 7:
			fmt.Printf("break %d\n", i)
			break Loop
		default:
			fmt.Printf("i=%d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Printf("break loop end i=%d\n", i)
}

func continueLoop() {
	var i = 0
Loop:
	for i = 0; i < 10; i++ {
		switch {
		case i > 7:
			fmt.Printf("continue %d\n", i)
			continue Loop
		default:
			fmt.Printf("i=%d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Printf("continue loop end i=%d\n", i)
}

func main() {
	breakLoop()
	continueLoop()
}

//==================================== END ======================================
