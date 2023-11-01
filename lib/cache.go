package lib

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/syndtr/goleveldb/leveldb"
	de "github.com/syndtr/goleveldb/leveldb/errors"
)

func OpenLocalDB() (*leveldb.DB, error) {
	dbDir := getDictDBDir()
	db, err := leveldb.OpenFile(dbDir, nil)
	if nil != err {
		return nil, err
	}
	return db, nil
}

func QueryLocalDB(key string, db *leveldb.DB) (*DictResult, error) {
	data, err := db.Get([]byte(key), nil)
	if de.ErrNotFound == err {
		// first query word always return NotFound
		return nil, nil
	}
	if nil != err {
		return nil, err
	}

	ret := DictResult{}
	if err := json.Unmarshal(data, &ret); nil != err {
		return nil, err
	}
	return &ret, nil
}

func ScanWords() (map[string]DictResult, error) {
	db, err := OpenLocalDB()
	if err != nil {
		color.Red("OpenLocalDb Fail! Cause: %s", err)
		return nil, err
	}

	defer db.Close()

	dict := map[string]DictResult{}
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		ret := DictResult{}
		if err := json.Unmarshal(iter.Value(), &ret); nil != err {
			return nil, err
		}

		dict[string(iter.Key())] = ret
	}

	defer iter.Release()
	return dict, nil
}

func DeleteWords(args []string) error {
	db, err := OpenLocalDB()
	if err != nil {
		color.Red("OpenLocalDb Fail! Cause: %s", err)
		return err
	}
	defer db.Close()

	if err := db.Delete([]byte(strings.Join(args, " ")), nil); err != nil {
		return err
	}

	return nil
}

func BackupCahceFiles() {
	dictDir := getDictDir()
	if _, err := os.Stat(dictDir); os.IsNotExist(err) {
		color.Red("Cannot find the DB path", err.Error())
		return
	}

	fileName := GetBakFileName()
	dstFile := filepath.Join(dictDir, fileName)

	result := Execute(dictDir, "tar", "-czvf", dstFile, "./db")
	if result {
		color.Green("Local DB backup to: %s", dstFile)
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
	dbDir := os.Getenv("YDICT_DB")
	xdgCacheDir := os.Getenv("XDG_CACHE_HOME")
	if dbDir == "" {
		if xdgCacheDir != "" {
			dbDir = filepath.Join(xdgCacheDir, "ydict")
		} else {
			dbDir = filepath.Join(os.Getenv("HOME"), ".cache/ydict")
		}
	}

	return dbDir
}

func getDictDBDir() string {
	dbDir := os.Getenv("YDICT_DB")
	xdgCacheDir := os.Getenv("XDG_CACHE_HOME")
	if dbDir == "" {
		if xdgCacheDir != "" {
			dbDir = filepath.Join(xdgCacheDir, "ydict")
		} else {
			dbDir = filepath.Join(os.Getenv("HOME"), ".cache/ydict")
		}
	}

	ydictDir := filepath.Join(dbDir, "db")
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
