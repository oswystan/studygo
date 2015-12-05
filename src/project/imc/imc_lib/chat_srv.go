//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: chat_srv.go
//     description:
//         created: 2015-12-05 21:57:08
//          author: wystan
//
//===============================================================================

package imc_lib

import (
	"log"
	"net"
)

type ImcServer struct {
}

func (s *ImcServer) Run() error {
	log.Println("imc server started...")
	ln, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Println("ERROR: %s", err)
		return err
	}
	defer ln.Close()

	for {
		con, err := ln.Accept()
		if err != nil {
			log.Println("ERROR: %s", err)
			return err
		}
		log.Println("new client arrived")
		con.Close()
	}

	return nil
}

func NewImcServer() *ImcServer {
	return &ImcServer{}
}

//==================================== END ======================================
