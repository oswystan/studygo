/*
 * *******************************************************************************
 *                     Copyright (c) 2015, wystan
 *       Filename:
 *
 *    Description:
 *
 *        Created:
 *       Revision:  none
 *
 *         Author:  wystan
 *
 * *******************************************************************************
 */

package commands

import (
	"fmt"
	"os"
)

type ShCommand struct {
	Run     func(subCmd *ShCommand, paras []string)
	CmdName string
	Usage   string
}

func RunCommand() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <command>\n", os.Args[0])
		return
	}

	if subCmd, ok := g_commands[os.Args[1]]; ok {
		subCmd.Run(subCmd, os.Args)
	} else {
		fmt.Printf("invalid command\n")
	}
}

var g_commands = make(map[string]*ShCommand)

func registerCommand(subCmd *ShCommand) {
	if subCmd == nil {
		return
	}
	g_commands[subCmd.CmdName] = subCmd
}

/**************************************** END ***********************************/
