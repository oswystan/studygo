//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: handler.go
//     description:
//         created: 2016-02-16 23:14:55
//          author: wystan
//
//===============================================================================

package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Relationship struct {
	Id    int64  `json:"id"`
	State string `json:"state"`
	Type  string `json:"type"`
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var ul []User
	var err error
	if ul, err = GetDB().ListUsers(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	for i := 0; i < len(ul); i++ {
		encoder.Encode(&ul[i])
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 128))
	if err != nil {
		log.Printf("ERROR: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body.Close()
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Printf("ERROR: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := GetDB()
	err = db.CreateUser(&user)
	if err != nil {
		log.Printf("ERROR: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		log.Printf("ERROR: %s", err)
		return
	}
}

func getRelationships(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("getRelationships\n"))
}

func createRelationship(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("createRelationship\n"))
}

//==================================== END ======================================
