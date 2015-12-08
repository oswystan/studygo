//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: server.go
//     description:
//         created: 2015-12-01 20:17:28
//          author: wystan
//
//===============================================================================

package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func handleConn(c net.Conn) error {
	defer c.Close()
	r := bufio.NewReader(c)
	w := bufio.NewWriter(c)

	for {
		str, err := r.ReadString(byte('\n'))
		if err != nil {
			log.Print(err)
			break
		}
		content := strings.TrimSuffix(str, "\n")
		if content == "bye" || content == "quit" {
			break
		}
		log.Printf("(%s)=>%s\n", c.RemoteAddr(), content)
		str = "hello " + str
		w.WriteString(str)
		w.Flush()
	}
	return nil
}

var c uint32 = 2000

func fun() error {
	return c
}

func main() {
	log.Printf("server started...\n")
	ln, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}

//==================================== END ======================================
