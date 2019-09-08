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
		err  error
		cmd  string
		arg2 string
		arg3 string
	)

	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	if len(os.Args) > 2 {
		arg2 = os.Args[2]
	}

	if len(os.Args) > 3 {
		arg3 = os.Args[3]
	}

	switch cmd {
	case "download":
		err = doCliDownload(arg2, arg3)
	case "build":
		toolchain := ""
		err = doCliBuild(arg2, toolchain)
	case "update":
		toolchain := ""
		err = doCliUpdate(toolchain)
	default:
		dumpUsage()
	}

	if err != nil {
		log.Fatal(err)
	}
}
