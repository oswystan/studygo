//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: imc_srv.go
//     description:
//         created: 2015-12-05 22:13:23
//          author: wystan
//
//===============================================================================

package main

import (
	"log"
	imc "project/imc/imservice"
)

func main() {
	srv := imc.NewImcServer()
	log.Fatal(srv.Run())
}

//==================================== END ======================================
