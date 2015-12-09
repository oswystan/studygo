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
	"log"
	"project/imc/imservice"
)

func main() {
	log.Println("client started ...")
	helper := imservice.NewHelper()
	helper.Connect("localhost", "8000")

	helper.RunCommand("login momo 123445")
	for {
		message := helper.PickMessage()
		if message == nil {
			break
		} else {
			log.Printf("%v", message)
		}
	}
	helper.RunCommand("logout momo")

	defer helper.Close()
}

//==================================== END ======================================
