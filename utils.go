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
	color.Cyan(logo, Version)
	color.Cyan("Usage:")
	color.Cyan("ydict <word(s) to query>        Query the word(s)")
	color.Cyan("ydict -v <word(s) to query>     Query with speech")
	color.Cyan("ydict -m <word(s) to query>     Query with more example sentences")
	color.Cyan("ydict -q <word(s) to query>     Query with quiet mode, don't show spinner")
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
	return runtime.GOOS == "darwin" || runtime.GOOS == "linux"
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

func parseArgs(args []string) ([]string, bool, bool, bool) {
	//match argument: -v or -m or -q
	var withVoice, withMore, isQuiet bool
	wordStartIndex := findWordStartIndex(args)
	paramArray := args[:wordStartIndex]
	if elementInStringArray(paramArray, "-m") {
		withMore = true
	}

	if elementInStringArray(paramArray, "-v") {
		withVoice = true
	}

	if elementInStringArray(paramArray, "-q") {
		isQuiet = true
	}

	return args[wordStartIndex:], withVoice, withMore, isQuiet
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
