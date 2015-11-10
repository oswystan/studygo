/*
 * *******************************************************************************
 *                     Copyright (c) 2015, wystan
 *       Filename:  main.go
 *
 *    Description:
 *
 *        Created:  2015-11-08 21:38:57
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
	"io"
	"os"
	"strings"
)

type LocaleReader struct {
	lang    string
	baseDir string
	desc    map[string]string
}

var localer = make(map[string]*LocaleReader)

func (l *LocaleReader) Get(key string) string {
	return l.desc[key]
}

func (l *LocaleReader) Init(lang string, dir string) error {
	dir += "/" + lang + ".ini"
	if l.desc == nil {
		l.desc = make(map[string]string)
	}
	f, err := os.Open(dir)
	if err != nil {
		fmt.Printf("fail to open file [%s]\n", err)
		return err
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		oneLine, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("fail to read file [%s]\n", err)
			break
		}

		strLine := string(oneLine)
		strLine = strings.TrimSpace(strLine)
		if len(strLine) == 0 || strLine[0] == '#' {
			continue
		}
		fields := strings.Split(strLine, "=")
		if len(fields) != 2 {
			continue
		}
		l.desc[strings.TrimSpace(fields[0])] = strings.TrimSpace(fields[1])
	}

	return nil
}

func main() {
	l := new(LocaleReader)
	l.Init("en", "language")
	localer["en"] = l
	l = new(LocaleReader)
	l.Init("zh", "language")
	localer["zh"] = l

	fmt.Printf("%s\n", localer["en"].Get("username"))
	fmt.Printf("%s\n", localer["zh"].Get("password"))
	l = nil
}

/**************************************** END ***********************************/
