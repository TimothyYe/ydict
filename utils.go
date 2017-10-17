package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"unicode"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

const (
	logo = `
██╗   ██╗██████╗ ██╗ ██████╗████████╗
╚██╗ ██╔╝██╔══██╗██║██╔════╝╚══██╔══╝
 ╚████╔╝ ██║  ██║██║██║        ██║   
  ╚██╔╝  ██║  ██║██║██║        ██║   
   ██║   ██████╔╝██║╚██████╗   ██║   
   ╚═╝   ╚═════╝ ╚═╝ ╚═════╝   ╚═╝   

YDict V0.9
https://github.com/TimothyYe/ydict

`
)

func displayUsage() {
	color.Cyan(logo)
	color.Cyan("Usage:")
	color.Cyan("ydict <word(s) to query>        Query the word(s)")
	color.Cyan("ydict <word(s) to query> -v     Query with speech")
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
	n := 0
	var withVoice, withMore bool
	for i := len(args) - 1; i > 0; i-- {
		if args[i] == "-v" {
			withVoice = true
			n++
		} else if args[i] == "-m" {
			withMore = true
			n++
		} else {
			break
		}
	}
	return args[1 : len(args)-n], withVoice, withMore
}
