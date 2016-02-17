//===============================================================================
//                      Copyright (C) 2016 wystan
//
//        filename: db.go
//     description:
//         created: 2016-02-16 23:44:24
//          author: wystan
//
//===============================================================================

package main

import "fmt"

type DB struct {
	connected bool
}

var db *DB = &DB{false}

var users = []User{
	{1, "bob", "user"},
	{2, "mike", "user"},
	{3, "jason", "user"},
}

var rs = []Relationship{
	{},
}

func GetDB() *DB {
	return db
}

func (db *DB) Connect(srv string, user string, pwd string) error {
	if db.connected {
		return fmt.Errorf("already connected")
	}

	db.connected = true
	return nil
}

func (db *DB) Disconnect() error {
	if !db.connected {
		return nil
	}

	db.connected = false
	return nil
}

func (db *DB) CreateUser(u *User) error {
	u.Id = (int64)(len(users) + 1)
	u.Type = "user"
	users = append(users, *u)
	return nil
}

func (db *DB) ListUsers() ([]User, error) {
	return users[0:], nil
}

func (db *DB) GetRelationships(id int64) ([]Relationship, error) {
	return nil, nil
}
func (db *DB) CreateRelationship(id1 int64, id2 int64, rs int) error {
	return nil
}
func (db *DB) UpdateRelationship(id1 int64, id2 int64, rs int) error {
	return nil
}
func (db *DB) GetRelationship(id1 int64, id2 int64) (int, error) {
	return 1, nil
}

//==================================== END ======================================
