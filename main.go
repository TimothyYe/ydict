package main

import (
	"os"
)

var (
	proxy string
)

func main() {
	//Check & load .env file
	loadEnv()

	if len(os.Args) == 1 {
		displayUsage()
		os.Exit(0)
	}

	if len(os.Args) == 2 && os.Args[1] == "-h" {
		displayUsage()
		os.Exit(0)
	}

	words, withVoice, withMore, isQuiet := parseArgs(os.Args[1:])
	query(words, withVoice, withMore, isQuiet, len(words) > 1)
}
