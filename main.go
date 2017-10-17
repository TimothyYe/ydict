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

	words, withVoice, withMore := parseArgs(os.Args)
	query(words, withVoice, withMore, len(words) > 1)
}
