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
)

var helpCmd = &ShCommand{
	Run:     doHelp,
	CmdName: "help",
	Usage:   "show this message",
}

func doHelp(subCmd *ShCommand, params []string) {
	fmt.Printf("\navaliable commands: \n")
	for _, cmd := range g_commands {
		fmt.Printf("\t%s\t\t:%s\n", cmd.CmdName, cmd.Usage)
	}
	fmt.Printf("\n")
}

func init() {
	registerCommand(helpCmd)
}

/**************************************** END ***********************************/
