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
	log.Printf("starting server...")
	db := GetDB()
	if err := db.Connect("socialdb", "pgtest", "123456"); err != nil {
		log.Printf("ERROR: %s", err)
		return
	}
	log.Printf("database connected")

	r := NewRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
	db.Disconnect()
}

//==================================== END ======================================
