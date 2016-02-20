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

import (
	"fmt"

	"gopkg.in/pg.v3"
)

const (
	NO_RELATIONSHIP = iota
	LIKED
	DISLIKED
	MATCHED
)

type Database struct {
	pg *pg.DB
}

var db *Database = &Database{nil}

var states = []string{
	"",
	"liked",
	"disliked",
	"matched",
}

func GetDB() *Database {
	return db
}

func postProcessRs(rs []Relationship, p int) {
	for i := 0; i < len(rs); i++ {
		r := &rs[i]
		if p == 1 {
			r.Id = r.Peer2
			r.State = states[r.Relation1]
		} else if p == 2 {
			r.Id = r.Peer1
			r.State = states[r.Relation2]
		}
		if r.Relation1 == LIKED && r.Relation2 == LIKED {
			r.State = states[MATCHED]
		}
	}

}

func (db *Database) checkUserId(id int64) ([]User, error) {
	var ul []User
	_, err := db.pg.Query(&ul, "select * from users where id=?", id)
	return ul, err
}

func (db *Database) checkUserName(name string) ([]User, error) {
	var ul []User
	_, err := db.pg.Query(&ul, "select * from users where name=?", name)
	return ul, err
}
func (db *Database) checkRs(id1 int64, id2 int64) ([]Relationship, error) {
	var rs []Relationship
	_, err := db.pg.Query(&rs, "select * from relationships where peer1=? and peer2=?",
		id1, id2)
	return rs, err
}
func (db *Database) clearUser() error {
	_, err := db.pg.Exec("delete from users; alter sequence users_id_seq restart with 1;")
	return err
}
func (db *Database) clearRs() error {
	_, err := db.pg.Exec("delete from relationships;")
	return err
}

func (db *Database) Connect(dbname string, user string, pwd string) error {
	if db.pg != nil {
		return fmt.Errorf("already connected")
	}
	opt := &pg.Options{}
	if len(user) != 0 {
		opt.User = user
		if len(pwd) != 0 {
			opt.Password = pwd
		}
	}
	opt.Database = dbname
	db.pg = pg.Connect(opt)

	if db.pg == nil {
		return fmt.Errorf("fail to connect database")
	}
	return nil
}

func (db *Database) Disconnect() error {
	if db.pg == nil {
		return nil
	}

	db.pg.Close()
	db.pg = nil
	return nil
}

func (db *Database) CreateUser(u *User) error {
	if db.pg == nil {
		return fmt.Errorf("database disconnected!")
	}

	// don't need to check whether the username is already existed,
	// because a unique index on users.name colomn.
	// but there is a bad effect: it will cost a squence number :-(
	// insert it into database
	_, err := db.pg.QueryOne(u, "insert into users(name) values(?) returning id;", u.Name)
	if err != nil {
		return err
	}
	u.Type = "user"
	return nil
}

func (db *Database) ListUsers() ([]User, error) {
	if db.pg == nil {
		return nil, fmt.Errorf("database disconnected!")
	}

	// TODO we should return data by page
	var ul []User
	_, err := db.pg.Query(&ul, "select id, name, 'user' as type from users;")
	if err != nil {
		return nil, err
	}

	return ul, nil
}

func (db *Database) GetRelationships(id int64) ([]Relationship, error) {
	var rs []Relationship
	sql := fmt.Sprintf("select *, 'relationship' as type from relationships where peer1=? and relation1 != 0;")
	_, err := db.pg.Query(&rs, sql, id)
	if err != nil {
		return rs, err
	}
	idx := len(rs)
	postProcessRs(rs, 1)

	sql = fmt.Sprintf(`select *, 'relationship' as type from relationships 
			where peer2=? and relation2 != 0;`)
	_, err = db.pg.Query(&rs, sql, id)
	if err != nil {
		return rs, err
	}
	postProcessRs(rs[idx:], 2)

	if len(rs) == 0 {
		return rs, fmt.Errorf("no relationships found")
	}
	return rs, nil
}
func (db *Database) CreateRelationship(id1 int64, id2 int64, state string) (*Relationship, error) {
	var rs int64
	switch state {
	case "liked":
		rs = LIKED
	case "disliked":
		rs = DISLIKED
	default:
		return nil, fmt.Errorf("invalid relationship %s", rs)
	}

	if id1 == id2 {
		return nil, fmt.Errorf("should use different id(%d)", id1)
	}

	r := &Relationship{}

	_, err := db.pg.QueryOne(r, "select * from create_rs(?,?,?);", id1, id2, rs)
	if err != nil {
		return nil, err
	}

	r.Id = id2
	r.State = state
	r.Type = "relationship"
	if r.Relation1 == LIKED && r.Relation2 == LIKED {
		r.State = states[MATCHED]
	}

	return r, err
}

//==================================== END ======================================
