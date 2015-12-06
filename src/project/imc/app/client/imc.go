//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: imc.go
//     description:
//         created: 2015-12-06 09:19:42
//          author: wystan
//
//===============================================================================

package main

import (
	"bufio"
	"log"
	"net"
	imc "project/imc/imc_lib"
)

func main() {
	log.Println("client started ...")
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	proxy := imc.NewImcProxy(reader, writer)

	proxy.DoLogin("wangyu", "1234")
	proxy.DoModifyInfo("wangyu", "nopasswd", "oswystan")
}

//==================================== END ======================================
