/*
 * *******************************************************************************
 *                     Copyright (c) 2015, wystan
 *       Filename:  main.go
 *
 *    Description:
 *
 *        Created:  2015-11-08 15:06:49
 *       Revision:  none
 *
 *         Author:  wystan
 *
 * *******************************************************************************
 */

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type UserInfo struct {
	id       int
	name     string
	password string
}

func (u UserInfo) String() string {
	return fmt.Sprintf("%d|%s|%s", u.id, u.name, u.password)
}

func main() {
	db, err := sql.Open("mysql", "root:1@tcp(127.0.0.1:3306)/epmp?charset=utf8")

	if err != nil {
		fmt.Printf("fail to open mysql database <%s>\n", err)
		os.Exit(1)
	}

	rows, err := db.Query("select id, username, password from user_base")
	if err != nil {
		fmt.Printf("fail to do select <%s>\n", err)
		os.Exit(1)
	}

	fmt.Printf("----------------\n")
	for rows.Next() {
		var oneUser UserInfo
		err = rows.Scan(&oneUser.id, &oneUser.name, &oneUser.password)
		fmt.Printf("%s\n", oneUser)
		fmt.Printf("----------------\n")
	}
	db.Close()
}

/**************************************** END ***********************************/
