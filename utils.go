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
	color.Cyan("ydict <word(s) to query> -v     Query with speech")
	color.Cyan("ydict <word(s) to query> -m     Query with more example sentences")
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

func parseArgs(args []string) ([]string, bool, bool) {
	//match argument: -v or -m
	var withVoice, withMore bool
	parameterStartIndex := findParamStartIndex(args)
	paramArray := args[parameterStartIndex:]
	if elementInStringArray(paramArray, "-m") {
		withMore = true
	}

	if elementInStringArray(paramArray, "-v") {
		withVoice = true
	}
	return args[1:parameterStartIndex], withVoice, withMore
}

func findParamStartIndex(args []string) int {
	// iter the args array, if an element is -m or -v, then  all of the latter elements must be parameter instead of words.
	for index, word := range args {
		if strings.HasPrefix(word, "-") && len(word) == 2 {
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
