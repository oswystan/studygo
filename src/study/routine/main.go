/*
 * *******************************************************************************
 *                     Copyright (c) 2015, wystan
 *       Filename:  main.go
 *
 *    Description:
 *
 *        Created:  2015-11-08 14:16:58
 *       Revision:  none
 *
 *         Author:  wystan
 *
 * *******************************************************************************
 */

package main

import (
	"fmt"
	"time"
)

//show how go routine works
func studyGoroutine() {
	var rst int
	done := make(chan int)
	defer close(done)
	gofunc := func() {
		var i int
		for i = 0; i < 5; i++ {
			fmt.Printf("go run\n")
			time.Sleep(1000 * time.Millisecond)
		}
		done <- i
		fmt.Printf("go done\n")
		done <- i
	}
	go gofunc()
	rst = <-done
	fmt.Printf("main done %d\n", rst)
}

/*
 * main entry
 */
func main() {
	studyGoroutine()
}

/**************************************** END ***********************************/
