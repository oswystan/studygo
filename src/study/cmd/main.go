/*
 * Copyright (C) wystudio
 * Demo code
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func processCmd(cmds []string) {
}

func ProcessOneFile(fileName string) (err error) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("fail to open the file[%s]\n", fileName)
		return
	}
	defer f.Close()

	buf := bufio.NewReader(f)

	var retError error

	// read all lines and process it
	for i := 1; ; i++ {
		oneline, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("fail to read data from file [%s]\n", fileName)
			retError = err
			break
		}
		strLine := string(oneline)

		strLine = strings.TrimSpace(strLine)
		fmt.Printf("%d [%s]\n", i, strLine)
	}

	return retError
}

func main() {
	// check the command line options
	argc := len(os.Args)
	if argc != 2 {
		fmt.Printf("usage:%s <filename>\n", os.Args[0])
		return
	}

	// first step: open file
	fileName := os.Args[1]
	_ = ProcessOneFile(fileName)

	fmt.Printf("all things done ...\n")
}

/*********************************** END ********************************************/
