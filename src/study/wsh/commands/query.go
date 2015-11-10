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

var queryCmd = &ShCommand{
	Run:     doQuery,
	CmdName: "query",
	Usage:   "query",
}

func doQuery(queryCmd *ShCommand, params []string) {
	fmt.Printf("do query now ...\n")

	fmt.Printf("done\n")
}

func init() {
	registerCommand(queryCmd)
}

/**************************************** END ***********************************/
