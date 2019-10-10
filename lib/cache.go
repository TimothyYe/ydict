package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"errors"

	"github.com/fatih/color"
	"github.com/syndtr/goleveldb/leveldb"
	de "github.com/syndtr/goleveldb/leveldb/errors"
)

type dictResult struct {
	WordString string

	PartOfSpeech string
	Meanings     []string
	Hints        [][]string
	Pronounce    string
	Result       string
	Sentences    [][]string

	AudioFilePath string
}

func OpenLocalDB() (*leveldb.DB, error) {
	dbDir := getDictDBDir()
	db, err := leveldb.OpenFile(dbDir, nil)
	if nil != err {
		return nil, err
	}
	return db, nil
}

func QueryLocalDB(key string, db *leveldb.DB) (*dictResult, error) {
	data, err := db.Get([]byte(key), nil)
	if de.ErrNotFound == err {
		// first query word always return NotFound
		return nil, nil
	}
	if nil != err {
		return nil, err
	}

	ret := dictResult{}
	if err := json.Unmarshal(data, &ret); nil != err {
		return nil, err
	}
	return &ret, nil
}

func (this dictResult) RemoveAudioFile() error {
	if len(this.AudioFilePath) == 0 {
		return nil
	}
	err := os.Remove(this.AudioFilePath) // clean up
	if nil != err && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func (this dictResult) SaveLocalDB(db *leveldb.DB) error {
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

func (this dictResult) Print(fromTag string, playCount int) {
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

func ClearCahceFiles() {
	tmpDir := getDictDir()
	err := os.RemoveAll(tmpDir)
	if nil != err && !os.IsNotExist(err) {
		color.Red("ClearCacheFile Fail! Cause: %s", err.Error())
	}
	color.Green("Clear Success! CacheDir: %s", tmpDir)
}

func getDictDir() string {
	tmpDir := os.TempDir()
	ydictDir := filepath.Join(tmpDir, "ydict")
	return ydictDir
}
func getDictDBDir() string {
	tmpDir := getDictDir()
	ydictDir := filepath.Join(tmpDir, "db")
	return ydictDir
}

func getDictAudioDir() string {
	tmpDir := getDictDir()
	ydictDir := filepath.Join(tmpDir, "audio")
	return ydictDir
}

func SaveVoiceFile(name string, body io.ReadCloser) (string, error) {
	ydictDir := getDictAudioDir()
	tmpfile, err := ioutil.TempFile(ydictDir, name)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", err
		}
		err = os.MkdirAll(ydictDir, 0700)
		if nil != err {
			return "", err
		}
		tmpfile, err = ioutil.TempFile(ydictDir, name)
		if nil != err {
			return "", err
		}
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}

	if _, err := tmpfile.Write(data); err != nil {
		return "", err
	}

	if err := tmpfile.Close(); err != nil {
		return "", err
	}

	aFile := tmpfile.Name()
	return aFile, err

}

func DoPlayFile(aFile string) error {
	cmd := exec.Command("mpg123", aFile)
	if _, err := exec.LookPath("mpv"); err == nil {
		// andoird termux only have mpv
		cmd = exec.Command("mpv", aFile)
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
