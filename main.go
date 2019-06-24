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

	if len(os.Args) == 2 && os.Args[1] == "-h" {
		displayUsage()
		os.Exit(0)
	}

	words, withVoice, withMore, isQuiet, withCache, clearCache := parseArgs(os.Args[1:])
	if clearCache {
		ClearCahceFiles()
		return
	}

	queryP := queryParam{}
	queryP.Words = words
	queryP.WordString = strings.Join(words, " ")
	queryP.WithMore = withMore
	queryP.WithCache = withCache
	queryP.IsQuiet = isQuiet
	queryP.IsMulti = (len(words) > 1)
	queryP.IsChinese = isChinese(queryP.WordString)
	queryP.WithVoice = withVoice
	if !isAvailableOS() {
		queryP.WithVoice = 0
	}
	queryP.DoQuery()
}
