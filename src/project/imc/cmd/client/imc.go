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

var commands = [...]string{
	"login wangyu 123456",
	"modify wangyu 1234 oswystan",
	"talk momo hello momo, how are you doing?",
	"logout momo",
}

func main() {
	log.Println("client started ...")
	helper := imservice.NewHelper()
	helper.Connect("localhost", "8000")

	var err error
	for i := 0; i < len(commands); i++ {
		if err = helper.RunCommand(commands[i]); err != nil {
			log.Printf("Fail to run command:%s[%s]", commands[i], err)
			break
		}
	}

	defer helper.Close()
	if err == nil {
		log.Printf("all command success")
	}
}

//==================================== END ======================================
