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
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Relationship struct {
	Id        int64  `json:"id"`
	State     string `json:"state"`
	Type      string `json:"type"`
	Peer1     int64  `json:"-"`
	Peer2     int64  `json:"-"`
	Relation1 int    `json:"-"`
	Relation2 int    `json:"-"`
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var ul []User
	var err error
	if ul, err = GetDB().ListUsers(); err != nil {
		log.Printf("ERROR: fail to get user list %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ul)
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

	// get userid from url
	vars := mux.Vars(r)
	uid, _ := strconv.ParseInt(vars["uid"], 10, 64)
	db := GetDB()
	rs, err := db.GetRelationships(uid)
	if err != nil || rs == nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERROR: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rs)
}

func createRelationship(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	id1, _ := strconv.ParseInt(vars["uid"], 10, 64)
	id2, _ := strconv.ParseInt(vars["otheruid"], 10, 64)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 128))
	if err != nil {
		log.Printf("ERROR: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body.Close()
	rs := &Relationship{}
	if err := json.Unmarshal(body, rs); err != nil {
		log.Printf("ERROR: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := GetDB()
	rs, err = db.CreateRelationship(id1, id2, rs.State)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERROR: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rs)
}

//==================================== END ======================================
