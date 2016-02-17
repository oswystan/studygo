//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: main.go
//     description:
//         created: 2016-02-16 19:10:38
//          author: wystan
//
//===============================================================================
package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRouter()
	log.Printf("starting server...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

//==================================== END ======================================
