//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: main.go
//     description:
//         created: 2015-12-05 20:55:10
//          author: wystan
//
//===============================================================================

package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"study/protobuf/example"
)

func main() {
	test := &example.Test{
		Label:  proto.String("f"),
		Number: proto.Int32(20),
	}

	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Printf("err\n")
		return
	}

	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		fmt.Printf("err 2\n")
		return
	}

	fmt.Printf("msg=%v\n", newTest)
	fmt.Printf("len=%d\n", len(data))
}

//==================================== END ======================================
