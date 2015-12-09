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

//=================================
// 1. label usage in for loop
//=================================
func breakLoop() {
	var i = 0
Loop:
	for i = 0; i < 10; i++ {
		switch {
		case i > 7:
			fmt.Printf("break %d\n", i)
			//break the for loop
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
			//continue the for loop
			continue Loop
		default:
			fmt.Printf("i=%d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Printf("continue loop end i=%d\n", i)
}

//=================================
// the internal of defer
//=================================
func deferUse() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("i=%d\n", i)
	}

	//output:
	//  i=4
	//  i=3
	//  i=2
	//  i=1
	//  i=0
	// defer function is pushed into a stack.
	// the last one is called first
	// but the value of i is evaluated when the program
	// run the defer statement.
}

func vardef() {
	a := 0

	// although a is defined and given a value
	// here if there a more than one variable defined
	// you can ALSO use := operator. so convenient.
	v, a := 10, 20
	fmt.Printf("v=%d,a=%d\n", v, a)
}

//=================================
// different receiver
//     - a value receiver will not change the value outside
//     - a pointer receiver will change the value outside
//=================================
type Int int

func (i Int) change() {
	i = i + 1
}
func (i *Int) pchange() {
	*i = *i + 2
}

func receiver() {
	var i Int = 1
	fmt.Printf("before change i = %d\n", i)

	// because change() is a value receiver
	// i will not change outside
	i.change()
	fmt.Printf("after change i = %d\n", i)

	// pointer receiver
	// i changed.
	i.pchange()
	fmt.Printf("after pchange i = %d\n", i)
}

//=================================
// how to know the variable types
//=================================
func vartypes() {
	var s interface{}

	// define a string array, size is depent on the value initialized
	s = [...]string{0: "zero", 1: "one"}

	// use %T to print the type of var
	fmt.Printf("%T\n", s)

	switch t := s.(type) {
	case string:
		fmt.Printf("string\n")
	case [2]string:
		fmt.Printf("[2]string\n")
	default:
		fmt.Printf("unknown type %T\n", t)
	}
}

//=================================
// embeded type study
//=================================
type TypeA struct {
}

func (a *TypeA) Run() {
	fmt.Printf("TypeA is run\n")
}

type TypeB struct {
}

func (b *TypeB) RunIt() {
	fmt.Printf("TypeB is run\n")
}

type TypeC struct {
	/* data */
	*TypeA
	*TypeB
}

func embededFunc() {
	c := &TypeC{}
	c.Run()
	c.RunIt()
}

//=================================
// visiability of member and method
//=================================

// onetype can only be used in the current package
// and its members and methods can only be used in
// the current package too.
type onetype struct {
	value int
}

func (t *onetype) change() {
	t.value = 0
}

func visiability() {
	t := &onetype{}
	t.value = 1
	t.change()
	fmt.Printf("t.value=%d\n", t.value)
}

func main() {
	//breakLoop()
	//continueLoop()
	//deferUse()
	//vardef()
	//vartypes()
	//receiver()
	//embededFunc()
	//visiability()
}

//==================================== END ======================================
