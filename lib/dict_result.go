package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/syndtr/goleveldb/leveldb"
)

type DictResult struct {
	WordString string

	PartOfSpeech string
	Meanings     []string
	Hints        [][]string
	Pronounce    string
	Result       string
	Sentences    [][]string

	AudioFilePath string
}

func (this DictResult) RemoveAudioFile() error {
	if len(this.AudioFilePath) == 0 {
		return nil
	}
	err := os.Remove(this.AudioFilePath) // clean up
	if nil != err && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func (this DictResult) SaveLocalDB(db *leveldb.DB) error {
	if nil == db {
		return errors.New("invalid DB")
	}
	data, err := json.Marshal(this)
	if nil != err {
		return err
	}
	key := this.WordString
	if err := db.Put([]byte(key), data, nil); nil != err {
		return err
	}
	return nil
}

func (this DictResult) Print(fromTag string, playCount int) {
	if this.PartOfSpeech != "" {
		fmt.Println()
		fmt.Printf("%14s ", color.MagentaString(this.PartOfSpeech))
	}
	if len(this.Meanings) > 0 {
		fmt.Println()
		fmt.Printf("%s", color.GreenString(strings.Join(this.Meanings, "; ")))
	}
	if len(this.Hints) > 0 {
		fmt.Println()
		wordString := this.WordString
		color.Blue("    word '%s' not found, do you mean?", wordString)
		fmt.Println()
		for _, guess := range this.Hints {
			color.Green("    %s", guess[0])
			color.Magenta("    %s", guess[1])
		}
		fmt.Println()
	}
	if len(this.Pronounce) > 0 {
		fmt.Println()
		color.Green("    %s", this.Pronounce)
	}
	if len(this.Result) > 0 {
		fmt.Println()
		color.Green(this.Result)
	}

	if len(this.Sentences) > 0 {
		fmt.Println()
		for i, sentence := range this.Sentences {
			color.Green(" %2d.%s", i+1, sentence[0])
			color.Magenta("    %s", sentence[1])
		}
		fmt.Println()
	}

	if len(this.AudioFilePath) > 0 {
		for i := 0; i < playCount; i++ {
			err := DoPlayFile(this.AudioFilePath)
			if nil != err {
				color.Red("PlayFile Fail! Cause: %s", err.Error())
			}
		}
	}

	if len(fromTag) > 0 {
		fmt.Println()
		color.Blue("    [ %s ] From %s", this.WordString, fromTag)
	}
}
