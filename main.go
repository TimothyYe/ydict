package main

import (
	"os"
	"strings"

	"github.com/TimothyYe/ydict/lib"
)

func main() {
	//Check & load .env file
	lib.LoadEnv()

	if len(os.Args) == 1 {
		lib.DisplayUsage()
		os.Exit(0)
	}

	if len(os.Args) == 2 && os.Args[1] == "-h" {
		lib.DisplayUsage()
		os.Exit(0)
	}

	words, withVoice, withMore, isQuiet, withCache, clearCache := lib.ParseArgs(os.Args[1:])
	if clearCache {
		lib.ClearCahceFiles()
		return
	}

	queryP := lib.QueryParam{}
	queryP.Words = words
	queryP.WordString = strings.Join(words, " ")
	queryP.WithMore = withMore
	queryP.WithCache = withCache
	queryP.IsQuiet = isQuiet
	queryP.IsMulti = (len(words) > 1)
	queryP.IsChinese = lib.IsChinese(queryP.WordString)
	queryP.WithVoice = withVoice
	if !lib.IsAvailableOS() {
		queryP.WithVoice = 0
	}
	queryP.DoQuery()
}
