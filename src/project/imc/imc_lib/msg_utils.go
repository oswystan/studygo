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

import "github.com/golang/protobuf/proto"

func newCmdAckCommon(cmd CMD_TYPE, status RET_CODE, desc string) *ImcCmd {
	oneCmd := &ImcCmd{
		CmdType: &cmd,
		AckCommon: &CmdAckCommon{
			Status:    &status,
			ErrorDesc: proto.String(desc),
		},
	}

	return oneCmd
}

//==================================== END ======================================
