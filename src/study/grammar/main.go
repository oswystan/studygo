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
	"container/list"
	"fmt"
	"sync"
	"time"
	"unsafe"
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
	fmt.Printf("TypeB is run it\n")
}

type TypeC struct {

	// if the following types have the same member function(s),
	// you need to give a variable name to tell the compiler which
	// function you want to call.
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

	// you can set the member value outside the member function
	// but you can not set the member value outside the current package.
	t.value = 1
	t.change()
	fmt.Printf("t.value=%d\n", t.value)
}

//=================================
// show how to know the size of each type
//=================================
func sizeof() {
	var a int = 0
	var c bool = true
	var d struct{} = struct{}{}
	var m map[int]string
	var s string
	fmt.Printf("sizeof(int)=%d\n", unsafe.Sizeof(a))
	fmt.Printf("sizeof(bool)=%d\n", unsafe.Sizeof(c))
	fmt.Printf("sizeof(struct{})=%d\n", unsafe.Sizeof(d))
	fmt.Printf("sizeof(map[int]string)=%d\n", unsafe.Sizeof(m))
	fmt.Printf("sizeof(string)=%d\n", unsafe.Sizeof(s))
}

//=================================
// slice
//		a slice = nil (without make)
//		after make, all entries in slice is initialized to 0
//=================================
func slice() {
	var a []int
	if a != nil {
		fmt.Printf("not nil\n")
	} else {
		fmt.Printf("nil\n")
	}

	a = make([]int, 10, 11)
	fmt.Printf("len(a)=%d cap(a)=%d\n", len(a), cap(a))
	for idx, v := range a {
		fmt.Printf("a[%d]=%d\n", idx, v)
	}

	// DON'T know whether it it correct to do so
	a = nil
}

//=================================
// channel
//		- select a nil channel will wait for ever.
//      - in 'case' it will read or write to the channel. so
//			if you want to receive the data from channel ,
//			you must put a variable on the left or right;
//      - len(ch) = data len which has been enqueued into channel
//      - cap(ch) = the total capacity of channel when it is created.
//
// about nil channel and closed channel
//		- a closed channel can still use len() to get the data len in it
//      - A send to a nil channel blocks forever
//		- A receive from a nil channel blocks forever
//		- A send to a closed channel panics
//		- A receive from a closed channel returns the zero value immediately
//=================================
func channel() {
	ch := make(chan int, 10)
	fmt.Printf("len(ch)=%d\n", len(ch))
	ch <- 3
	fmt.Printf("after push 3 into it, len(ch)=%d cap(ch)=%d\n", len(ch), cap(ch))
	select {
	//even there is no receiver here,
	//the data will always be dequeued from channel
	//so len(ch)=0 HERE after case.
	case <-ch:
		fmt.Printf("len(ch)=%d\n", len(ch))
	}
	close(ch)
	if ch == nil {
		fmt.Printf("a closed channel = nil\n")
	} else {
		fmt.Printf("a closed channel != nil\n")
	}

	//close channel
	fmt.Printf("close channel: \n")
	ch2 := make(chan bool, 2)
	ch2 <- true
	ch2 <- true
	close(ch2)
	fmt.Printf("len(ch2)=%d\n", len(ch2))
	for i := 0; i < cap(ch2)+1; i++ {
		v, ok := <-ch2
		fmt.Println(v, ok)
	}
}

func closechstruct() {
	finish := make(chan struct{})
	var done sync.WaitGroup
	done.Add(1)
	go func() {
		select {
		// a VERY USEFUL way to do timeout
		case <-time.After(1 * time.Hour):
		case <-finish:
		}
		done.Done()
	}()
	t0 := time.Now()
	close(finish)
	done.Wait()
	fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0))
}

//==================================================
// env GODEBUG=gctrace=1,schedtrace=1000 ./grammar
//
// even one struct{xxx} nested into another struct{xxx} the gc can
// collect the memory too.
//==================================================
func memory() {
	type stdata struct {
		data [64 * 1024]byte
	}
	type stmemory struct {
		used int
		data *stdata
	}

	list := list.New()
	i := 0
	for {
		c := new(stmemory)
		d := new(stdata)
		c.data = d
		list.PushBack(c)
		time.Sleep(10 * time.Millisecond)
		if c == nil {
			break
		}
		i++
		if i%1024 == 0 {
			i = 0
			//this will cause gc to collect memory to the minimal size
			fmt.Printf("do list init\n")
			list.Init()
		}
	}
}

func main() {
	//breakLoop()
	//continueLoop()
	//deferUse()
	//vardef()
	//receiver()
	//vartypes()
	//embededFunc()
	//visiability()
	sizeof()
	//slice()
	//channel()
	//closechstruct()
	//memory()
}

//==================================== END ======================================
