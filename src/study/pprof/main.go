//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: main.go
//     description:
//         created: 2015-12-10 10:53:03
//          author: wystan
//
//===============================================================================

package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

type people struct {
	/* data */
	age  [1024]int
	name [1024]int
	size [1024]int
}

func doP(p []*people) {
	for i := 0; i < len(p); i++ {
		p[i].age[1023] = 111
	}

}

func heap() {
	f, err := os.Create("profile")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer f.Close()

	//==================================
	var p [1024]*people
	for i := 0; i < len(p); i++ {
		p[i] = &people{}
		p[i].age[1023] = 2
	}
	doP(p[:])
	for i := 0; i < len(p); i++ {
		p[i] = nil
	}

	var p2 [4 * 1024]*people
	for i := 0; i < len(p2); i++ {
		p2[i] = &people{}
		p2[i].age[1023] = 1
	}
	time.Sleep(3 * time.Second)

	//==================================
	profile := pprof.Lookup("heap")
	if profile != nil {
		profile.WriteTo(f, 1)
	}
}

func httpprof() {
	go func() {
		http.ListenAndServe(":7000", nil)
	}()
}

func profalloc() {
	var sl []*people = make([]*people, 1024, 1024)
	for {
		p := new(people)
		p.age[0] = 100
		p.age[1023] = 100
		time.Sleep(1 * time.Millisecond)
		sl = append(sl, p)
	}
}

type chType struct {
	/* data */
	data [256 * 1024]int
}

func profchannel() {

	var ch = make(chan *chType, 1024*1024*40)
	for i := 0; i < 1024; i++ {
		a := new(chType)
		a.data[1024] = 111
		ch <- a
	}
	for i := 0; i < 1024; i++ {
		p := <-ch
		if p != nil {
			p = nil
		}
	}

	fmt.Printf("sleep\n")
	runtime.GC()
	time.Sleep(100 * time.Second)
}

// for server:
//		go tool pprof -text $app http://localhost:$port/debug/pprof/heap
// for app client:
//		go tool pprof $app $profile

func main() {
	fmt.Printf("started ...\n")
	//httpprof()

	var i int = 10
	fmt.Printf("%p\n", &i)
	p := new(int)
	fmt.Printf("%p\n", p)
}

//==================================== END ======================================
