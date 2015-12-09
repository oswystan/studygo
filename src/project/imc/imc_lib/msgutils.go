//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: msg_utils.go
//     description:
//         created: 2015-12-06 18:13:09
//          author: wystan
//
//===============================================================================

package imc_lib

import (
	"bufio"
	"encoding/binary"
	"log"
)

const MAX_DATA_SIZE = 1024

func readMsgData(r *bufio.Reader) ([]byte, error) {
	var length int32 = 0

	err := binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}

	if length <= 0 || length > MAX_DATA_SIZE {
		log.Printf("invalid length=%d should be in (0,%d]", length, MAX_DATA_SIZE)
		panic("invalid msg length")
	}

	bs := make([]byte, length)
	_, err = r.Read(bs)
	return bs, err
}

func writeMsgData(w *bufio.Writer, data []byte) error {
	var length int32 = int32(len(data))
	if length > MAX_DATA_SIZE {
		log.Printf("invalid send msg size=%d", length)
		panic("too large msg data")
	}

	err := binary.Write(w, binary.BigEndian, length)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	w.Flush()

	return nil
}

//==================================== END ======================================
