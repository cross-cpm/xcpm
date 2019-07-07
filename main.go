package main

import (
	"fmt"
	"log"
	"os"
)

func dumpUsage() {
	fmt.Print(`usage: xcpm <command> <args>

commands:
   init       create packae file
   update     update dependency

For additional information, see https://github.com/cross-cpm/xcpm.git
`)
}

func main() {
	//log.SetFlags(0)

	var (
		cmd    string
		subcmd string
	)

	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	if len(os.Args) > 2 {
		subcmd = os.Args[2]
	}

	switch cmd {
	case "init":
		log.Println(cmd, subcmd)
	case "update":
		doCliUpdate()
	default:
		dumpUsage()
	}
}
