//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: db_test.go
//     description:
//         created: 2016-02-17 12:55:03
//          author: wystan
//
//===============================================================================
package main

import (
	"fmt"
	"testing"
	"time"
)

var dbname = "socialdb"
var dbuser = "pgtest"
var dbpwd = "123456"

func TestDBConnection(t *testing.T) {
	db := GetDB()
	err := db.Connect(dbname, dbuser, dbpwd)
	if err != nil {
		t.Errorf("%s", err)
	}
	db.Disconnect()
}

func TestCreateUser(t *testing.T) {
	db := GetDB()
	err := db.Connect(dbname, dbuser, dbpwd)
	if err != nil {
		t.Errorf("%s", err)
	}
	err = db.clearUser()
	if err != nil {
		t.Errorf("%s", err)
	}

	user := fmt.Sprintf("usr-%d", time.Now().UnixNano())
	usr := &User{Name: user}
	err = db.CreateUser(usr)
	if err != nil {
		t.Errorf("%s", err)
	}

	user = fmt.Sprintf("usr-%d", time.Now().UnixNano())
	usr = &User{Name: user}
	err = db.CreateUser(usr)
	if err != nil {
		t.Errorf("%s", err)
	}

	user = fmt.Sprintf("usr-%d", time.Now().UnixNano())
	usr = &User{Name: user}
	err = db.CreateUser(usr)
	if err != nil {
		t.Errorf("%s", err)
	}

	db.Disconnect()
}

func TestCreateRelationship(t *testing.T) {
	db := GetDB()
	err := db.Connect(dbname, dbuser, dbpwd)
	if err != nil {
		t.Errorf("%s", err)
	}
	err = db.clearRs()
	if err != nil {
		t.Errorf("%s", err)
	}

	_, err = db.CreateRelationship(1, 2, "liked")
	if err != nil {
		t.Errorf("%s", err)
	}

	_, err = db.CreateRelationship(2, 1, "disliked")
	if err != nil {
		t.Errorf("%s", err)
	}

	_, err = db.CreateRelationship(1, 1, "liked")
	if err == nil {
		t.Errorf("should check different id")
	}

	_, err = db.CreateRelationship(-1, 1, "liked")
	if err == nil {
		t.Errorf("should check the user id")
	}

	_, err = db.CreateRelationship(1, 2, "like")
	if err == nil {
		t.Errorf("should check invalid state")
	}
	db.Disconnect()

	_, err = db.CreateRelationship(1, 2, "matched")
	if err == nil {
		t.Errorf("should check invalid state")
	}

	db.Disconnect()
}

func TestGetRelationship(t *testing.T) {
	db := GetDB()
	err := db.Connect(dbname, dbuser, dbpwd)
	if err != nil {
		t.Errorf("%s", err)
	}
	if err = db.clearRs(); err != nil {
		t.Errorf("%s", err)
	}

	_, err = db.CreateRelationship(1, 2, "liked")
	if err != nil {
		t.Errorf("%s", err)
	}
	rs, err := db.GetRelationships(1)
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(rs) != 1 {
		t.Errorf("only one relationship should be returned, but %d returned", len(rs))
	}
	if rs[0].State != "liked" {
		t.Errorf("relationship should be liked but [%s]", rs[0].State)
	}

	_, err = db.CreateRelationship(2, 1, "liked")
	if err != nil {
		t.Errorf("%s", err)
	}
	rs, err = db.GetRelationships(1)
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(rs) != 1 {
		t.Errorf("only one relationship should be returned, but %d returned", len(rs))
	}
	if rs[0].State != "matched" {
		t.Errorf("relationship should be liked but [%s]", rs[0].State)
	}

	_, err = db.CreateRelationship(2, 1, "disliked")
	if err != nil {
		t.Errorf("%s", err)
	}
	rs, err = db.GetRelationships(2)
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(rs) != 1 {
		t.Errorf("only one relationship should be returned, but %d returned", len(rs))
	}
	if rs[0].State != "disliked" {
		t.Errorf("relationship should be disliked but [%s]", rs[0].State)
	}

	_, err = db.CreateRelationship(1, 3, "disliked")
	if err != nil {
		t.Errorf("%s", err)
	}
	rs, err = db.GetRelationships(1)
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(rs) != 2 {
		t.Errorf("only one relationship should be returned, but %d returned", len(rs))
	}
	if rs[0].State != "liked" || rs[1].State != "disliked" {
		t.Errorf("need relationships [liked, disliked] but [%s, %s]", rs[0].State, rs[1].State)
	}
	db.Disconnect()
}

func TestXXX(t *testing.T) {
	db := GetDB()
	err := db.Connect(dbname, dbuser, dbpwd)
	if err != nil {
		t.Errorf("%s", err)
	}

	db.clearUser()
	db.clearRs()
	db.Disconnect()
}

//==================================== END ======================================
