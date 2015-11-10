/*
 * *******************************************************************************
 *                     Copyright (c) 2015, wystan
 *       Filename:  main.go
 *
 *    Description:
 *
 *        Created:  2015-11-08 14:19:52
 *       Revision:  none
 *
 *         Author:  wystan
 *
 * *******************************************************************************
 */

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type UserInfo struct {
	/* data */
	Name string
	Age  int
	OS   string
}

func (u *UserInfo) Init() error {
	u.Name = "wangyu"
	u.Age = 1000
	f, err := os.Open("/etc/issue.net")
	if err != nil {
		fmt.Printf("fail to open [%s]\n", err)
		return err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	oneline, _, err := reader.ReadLine()

	u.OS = string(oneline)
	return nil
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	u := new(UserInfo)
	u.Init()
	t := template.New("first")
	s1, _ := t.ParseFiles("index.html")
	s1.ExecuteTemplate(w, "index.html", u)
}

func main() {
	http.HandleFunc("/hello", sayHello)
	//err := http.ListenAndServe(":8080", nil)
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		fmt.Printf("fail to start http server %s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

/**************************************** END ***********************************/
