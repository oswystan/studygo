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
		r.Type = "relationship"
	}

}
func postProcessUser(ul []User) {
	for i := 0; i < len(ul); i++ {
		ul[i].Type = "user"
	}

}

func (db *Database) checkUserId(id int64) ([]User, error) {
	var ul []User
	sql := fmt.Sprintf("select * from users where id=%d", id)
	_, err := db.pg.Query(&ul, sql)
	return ul, err
}

func (db *Database) checkUserName(name string) ([]User, error) {
	var ul []User
	sql := fmt.Sprintf("select * from users where name=%s", name)
	_, err := db.pg.Query(&ul, sql)
	return ul, err
}
func (db *Database) checkRs(id1 int64, id2 int64) ([]Relationship, error) {
	var rs []Relationship
	sql := fmt.Sprintf("select * from relationships where peer1=%d and peer2=%d", id1, id2)
	_, err := db.pg.Query(&rs, sql)
	return rs, err
}
func (db *Database) clearUser() error {
	sql := "delete from users; alter sequence users_id_seq restart with 1;"
	_, err := db.pg.Exec(sql)
	return err
}
func (db *Database) clearRs() error {
	sql := "delete from relationships;"
	_, err := db.pg.Exec(sql)
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
	sql := fmt.Sprintf("insert into users(name) values('%s') returning id;", u.Name)
	_, err := db.pg.QueryOne(u, sql)
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

	// TODO we should get page of data
	var ul []User
	sql := "select id, name from users;"
	_, err := db.pg.Query(&ul, sql)
	if err != nil {
		return nil, err
	}
	postProcessUser(ul)

	return ul, nil
}

func (db *Database) GetRelationships(id int64) ([]Relationship, error) {
	var rs []Relationship
	sql := fmt.Sprintf("select * from relationships where peer1=%d and relation1 != 0;", id)
	_, err := db.pg.Query(&rs, sql)
	if err != nil {
		return rs, err
	}
	idx := len(rs)
	postProcessRs(rs, 1)

	sql = fmt.Sprintf("select * from relationships where peer2=%d and relation2 != 0;", id)
	_, err = db.pg.Query(&rs, sql)
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
	var p1, p2, r1, r2, rs int64
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

	// check whether the uid is existed
	ul, err := db.checkUserId(id1)
	if err != nil || len(ul) == 0 {
		return nil, fmt.Errorf("No user found for id %d", id1)
	}
	ul, err = db.checkUserId(id2)
	if err != nil || len(ul) == 0 {
		return nil, fmt.Errorf("No user found for id %d", id2)
	}

	if id1 < id2 {
		p1 = id1
		p2 = id2
		r1 = rs
	} else {
		p1 = id2
		p2 = id1
		r2 = rs
	}
	relation, err := db.checkRs(p1, p2)
	if err != nil {
		return nil, fmt.Errorf("fail to check relationship in database")
	}

	r := &Relationship{}
	if len(relation) == 0 {
		sql := fmt.Sprintf("insert into relationships values(%d, %d, %d, %d) returning peer1, peer2, relation1, relation2;",
			p1, p2, r1, r2)
		_, err = db.pg.QueryOne(r, sql)
	} else {
		sql := fmt.Sprintf("update relationships set relation1=%d where peer1=%d and peer2=%d returning relation1, relation2;",
			r1, p1, p2)
		if r2 > 0 {
			sql = fmt.Sprintf("update relationships set relation2=%d where peer1=%d and peer2=%d returning relation1, relation2;",
				r2, p1, p2)
		}
		_, err = db.pg.QueryOne(r, sql)
	}
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
