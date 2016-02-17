//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: router.go
//     description:
//         created: 2016-02-16 23:01:58
//          author: wystan
//
//===============================================================================

package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/users", listUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{uid:[0-9]+}/relationships", getRelationships).Methods("GET")
	r.HandleFunc("/users/{uid:[0-9]+}/relationships/{otheruid:[0-9]+}", createRelationship).Methods("PUT")

	return r
}

//==================================== END ======================================
