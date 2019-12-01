package lib

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"unicode"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

// GetBakFileName generates a backup file name by current date and time
func GetBakFileName() string {
	return fmt.Sprintf("ydict-%s.tar.gz", time.Now().Format("20060102150405"))
}

// Execute executes shell commands with arguments
func Execute(workDir, script string, args ...string) bool {
	cmd := exec.Command(script, args...)

	if workDir != "" {
		cmd.Dir = workDir
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		color.Red("%s", err.Error())
		return false
	}

	return true
}

func IsChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func IsAvailableOS() bool {
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

func LoadEnv() {
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

func ParseArgs(args []string) (
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
