package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var (
	//Version of ydict
	Version = "0.1"
	logo    = `
██╗   ██╗██████╗ ██╗ ██████╗████████╗
╚██╗ ██╔╝██╔══██╗██║██╔════╝╚══██╔══╝
 ╚████╔╝ ██║  ██║██║██║        ██║   
  ╚██╔╝  ██║  ██║██║██║        ██║   
   ██║   ██████╔╝██║╚██████╗   ██║   
   ╚═╝   ╚═════╝ ╚═╝ ╚═════╝   ╚═╝   

YDict V%s
https://github.com/TimothyYe/ydict

`
)

func displayUsage() {
	logo = ""
	color.Cyan(logo, Version)
	color.Cyan("Usage:")
	color.Cyan("ydict <word(s) to query>        Query the word(s)")
	color.Cyan("ydict -v <word(s) to query>     Query with speech")
	color.Cyan("ydict -m <word(s) to query>     Query with more example sentences")
	color.Cyan("ydict -q <word(s) to query>     Query with quiet mode, don't show spinner")
	color.Cyan("ydict -c <word(s) to query>     Query with local cache")
	color.Cyan("ydict -clear                    Clear local cache")
	color.Cyan("ydict -h                        For help")
}

func isChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func isAvailableOS() bool {
	switch runtime.GOOS {
	case "android":
		return true
	case "darwin":
		return true
	case "linux":
		return true
	}
	return false
}

func getExecutePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return filepath.Dir(ex)
}

func loadEnv() {
	exPath := getExecutePath()
	envPath := fmt.Sprintf("%s/.env", exPath)

	// if .env file doesn't exist, just return
	if _, err := os.Stat(fmt.Sprintf("%s/.env", exPath)); os.IsNotExist(err) {
		return
	}

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	proxy = os.Getenv("SOCKS5")
}

func parseArgs(args []string) (
	words []string,
	withVoice int,
	withMore bool,
	isQuiet bool,
	withCache bool,
	clearCache bool,
) {
	//match argument: -v or -m or -q
	wordStartIndex := findWordStartIndex(args)
	paramArray := args[:wordStartIndex]
	if elementInStringArray(paramArray, "-m") {
		withMore = true
	}

	withVoice = countInStringArray(paramArray, "-v")

	if elementInStringArray(paramArray, "-q") {
		isQuiet = true
	}

	if elementInStringArray(paramArray, "-c") {
		withCache = true
	}

	if elementInStringArray(paramArray, "-clear") {
		clearCache = true
	}

	return args[wordStartIndex:], withVoice, withMore, isQuiet, withCache, clearCache
}

func findWordStartIndex(args []string) int {
	// iter the args array, if an element is -m or -v or -q,
	// then all of the latter elements must be parameter instead of words.
	for index, word := range args {
		if !strings.HasPrefix(word, "-") {
			return index
		}
	}
	return len(args)
}

func elementInStringArray(stringArray []string, element string) bool {
	for _, word := range stringArray {
		if word == element {
			return true
		}
	}
	return false
}

func countInStringArray(stringArray []string, element string) int {
	count := 0
	for _, word := range stringArray {
		if word == element {
			count++
		}
	}
	return count
}
