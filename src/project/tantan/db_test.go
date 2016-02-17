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

import "testing"

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

	usr := &User{Name: "mike"}
	err = db.CreateUser(usr)
	if err != nil {
		t.Errorf("%s", err)
	}
	db.Disconnect()
}

func Testxxx(t *testing.T) {
}

//==================================== END ======================================
