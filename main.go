package main

import (
	"os"
	"strings"
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

	words := strings.Join(os.Args[1:], " ")
	query(words)
}
