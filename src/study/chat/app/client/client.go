//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: client.go
//     description:
//         created: 2015-12-01 20:17:54
//          author: wystan
//
//===============================================================================

package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	log.Print("client started ...")
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// write some strings to the server
	// and receive the response
	w := bufio.NewWriter(conn)
	w.WriteString("client run\n")
	w.Flush()

	r := bufio.NewReader(conn)
	str, err := r.ReadString(byte('\n'))
	if err != nil {
		log.Print(err)
	} else {
		log.Print(str)
	}
	w.WriteString("quit\n")
	w.Flush()
}

//==================================== END ======================================
