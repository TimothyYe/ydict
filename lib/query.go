package lib

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/gen2brain/beeep"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/xwjdsh/fy"
	proxier "golang.org/x/net/proxy"
)

var (
	proxy string
)

type QueryParam struct {
	Words      []string
	WordString string

	IsSentence bool
	IsQuiet    bool
	IsChinese  bool
	IsMulti    bool
	WithMore   bool
	WithCache  bool
	WithVoice  int
}

func (this QueryParam) DoQuery() {
	if len(this.WordString) == 0 {
		return
	}

	var dbCache *leveldb.DB = nil
	if !this.IsSentence && this.WithCache {
		if ldb, err := OpenLocalDB(); nil != err {
			color.Red("OpenLocalDb Fail! Cause: %s", err)
		} else {
			dbCache = ldb
		}
		if nil != dbCache {
			defer dbCache.Close()
			key := this.WordString
			ret, err := QueryLocalDB(key, dbCache)
			if nil != err {
				color.Red("QueryLocalDB Fail! Cause: %s", err)
			}
			if nil != ret {
				ret.Print("cache", this.WithVoice)
				return
			}
		}
	}

	//Init spinner
	var s *spinner.Spinner

	// disable the spinner on windows system
	if runtime.GOOS == "windows" {
		this.IsQuiet = true
	}

	if !this.IsQuiet {
		s = spinner.New(spinner.CharSets[35], 100*time.Millisecond)
		s.Prefix = "Querying... "
		if err := s.Color("green"); err != nil {
			color.Red("Failed to set color for spinner")
			os.Exit(1)
		}
		s.Start()
	}

	var (
		doc           *goquery.Document
		docMore       *goquery.Document
		audioFilePath string

		// for sentence
		sentenceResponse *fy.Response
	)

	if this.IsSentence {
		req := fy.Request{
			FromLang: fy.Chinese,
			ToLang:   fy.English,
			Text:     this.WordString,
		}
		if !this.IsChinese {
			req.FromLang, req.ToLang = req.ToLang, req.FromLang
		}
		sentenceResponse = fy.YoudaoTranslate(context.Background(), req)
	} else {
		doc, docMore, audioFilePath = this.ReqWeb()
	}

	if !this.IsQuiet {
		s.Stop()
	}

	if this.IsSentence {
		if err := sentenceResponse.Err; err != nil {
			color.Red("Some Thing Wrong! Cause: %s", err)
			os.Exit(1)
		}

		fmt.Println()
		color.Green("    %s", sentenceResponse.Result)
		fmt.Println()
		return
	}

	ret := this.ParseWeb(doc, docMore)
	ret.AudioFilePath = audioFilePath

	ret.Print("", this.WithVoice)

	if this.WithCache {
		err := ret.SaveLocalDB(dbCache)
		if nil != err {
			color.Red("Some Thing Wrong! Cause: %s", err)
		}
	} else {
		err := ret.RemoveAudioFile()
		if nil != err {
			color.Red("Some Thing Wrong! Cause: %s", err)
		}
	}
}

func (this QueryParam) ReqWeb() (
	doc *goquery.Document,
	docMore *goquery.Document,
	audioFilePath string,
) {
	var voiceBody io.ReadCloser

	urlMore := "http://dict.youdao.com/example/blng/eng/%s"
	urlVoice := "https://dict.youdao.com/dictvoice?audio=%s&type=2"
	urlQuery := ""
	if this.IsChinese {
		urlQuery = "https://dict.youdao.com/w/eng/%s"
	} else {
		urlQuery = "https://dict.youdao.com/w/%s"
	}

	queryString := strings.Join(this.Words, " ")
	voiceString := strings.Join(this.Words, "+")
	moreString := strings.Join(this.Words, "_")
	queryURL := fmt.Sprintf(urlQuery, queryString)
	voiceURL := fmt.Sprintf(urlVoice, voiceString)
	moreURL := fmt.Sprintf(urlMore, moreString)

	if proxy != "" {
		client := &http.Client{}
		dialer, err := proxier.SOCKS5("tcp", proxy, nil, proxier.Direct)

		if err != nil {
			color.Red("Can't connect to the proxy: %s", err)
			os.Exit(1)
		}

		httpTransport := &http.Transport{}
		client.Transport = httpTransport
		httpTransport.Dial = dialer.Dial

		resp, err := client.Get(queryURL)
		if err != nil {
			color.Red("Query failed with err: %s", err.Error())
			os.Exit(1)
		}
		doc, _ = goquery.NewDocumentFromReader(resp.Body)

		if this.WithVoice > 0 {
			if resp, err := client.Get(voiceURL); err == nil {
				voiceBody = resp.Body
				audioFilePath, err = SaveVoiceFile(this.WordString, voiceBody)
				if nil != err {
					color.Red("SaveVoiceFile failed with err: %s", err.Error())
				}
			}
		}
	} else {

		resp, err := http.Get(queryURL)
		if err != nil {
			color.Red("Query failed with err: %s", err.Error())
			os.Exit(1)
		}
		doc, _ = goquery.NewDocumentFromReader(resp.Body)

		if this.WithVoice > 0 {
			if resp, err := http.Get(voiceURL); err == nil {
				voiceBody = resp.Body
				audioFilePath, err = SaveVoiceFile(this.WordString, voiceBody)
				if nil != err {
					color.Red("SaveVoiceFile failed with err: %s", err.Error())
				}
			}
		}

	}

	if this.WithMore {
		if resp, err := http.Get(moreURL); err != nil {
			color.Red("Query failed with err: %s", err.Error())
			os.Exit(1)
		} else {
			docMore, _ = goquery.NewDocumentFromReader(resp.Body)
		}
	}

	return doc, docMore, audioFilePath
}

func (this QueryParam) ParseWeb(doc, docMore *goquery.Document) DictResult {
	ret := DictResult{}
	ret.WordString = this.WordString
	if this.IsChinese {
		// Find the result
		doc.Find(".trans-container > ul > p").Each(func(i int, s *goquery.Selection) {
			ret.PartOfSpeech = s.Children().Not(".contentTitle").Text()

			meanings := []string{}
			s.Find(".contentTitle > .search-js").Each(func(ii int, ss *goquery.Selection) {
				meanings = append(meanings, ss.Text())
			})
			ret.Meanings = meanings
		})
	} else {
		// Check for typos
		if hint := getHint(doc); hint != nil {
			ret.Hints = hint
			return ret
		}

		// Find the pronounce
		if !this.IsMulti {
			ret.Pronounce = getPronounce(doc)
		}

		// Find the result
		ret.Result = doc.Find("div#phrsListTab > div.trans-container > ul").Text()
	}

	// Show examples
	if nil != docMore {
		ret.Sentences = this.getSentences(docMore)
	} else {
		ret.Sentences = this.getSentences(doc)
	}

	return ret
}

func (this QueryParam) getSentences(doc *goquery.Document) [][]string {
	isChinese := this.IsChinese
	result := [][]string{}
	doc.Find("#bilingual ul li").Each(func(_ int, s *goquery.Selection) {
		r := []string{}
		s.Children().Each(func(ii int, ss *goquery.Selection) {
			// Ignore source
			if ii == 2 {
				return
			}
			var sentence string
			ss.Children().Each(func(iii int, sss *goquery.Selection) {
				if text := strings.TrimSpace(sss.Text()); text != "" {
					addSpace := (ii == 1 && isChinese) || (ii == 0 && !isChinese) && iii != 0 && text != "."
					if addSpace {
						text = " " + text
					}
					sentence += text
				}
			})
			r = append(r, sentence)
		})
		if len(r) == 2 {
			result = append(result, r)
		}
	})
	return result
}

func ListWords(withPlay int) {
	if dict, err := ScanWords(); err != nil {
		color.Red("  Failed to scan words from the cache.")
	} else {
		for k, v := range dict {
			fmt.Printf("%s \t %s", color.CyanString("%s", k), color.GreenString("%s", v.Result))
		}
	}
}

func DisplayWords(withPlay int) {
	if dict, err := ScanWords(); err != nil {
		color.Red("  Failed to scan words from the cache.")
	} else {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		go func() {
			<-c
			os.Exit(0)
		}()

		for {
			for k, v := range dict {
				message := k
				if len(v.Pronounce) > 0 {
					message = fmt.Sprintf("%s\r\n    %s\r\n", message, v.Pronounce)
				}
				// if len(v.Result) > 0 {
				// 	message = fmt.Sprintf("\r\n%s%s", message, v.Result)
				// }

				if err := beeep.Notify("YDict", message, ""); err != nil {
					color.Red(err.Error())
				}

				time.Sleep(time.Second * time.Duration(withPlay))
			}
		}
	}
}

func getPronounce(doc *goquery.Document) string {
	var pronounce string
	doc.Find("div.baav > span.pronounce").Each(func(i int, s *goquery.Selection) {

		if i == 0 {
			p := fmt.Sprintf("英: %s    ", s.Find("span.phonetic").Text())
			pronounce += p
		}

		if i == 1 {
			p := fmt.Sprintf("美: %s", s.Find("span.phonetic").Text())
			pronounce += p
		}
	})

	return pronounce
}

func getHint(doc *goquery.Document) [][]string {
	typos := doc.Find(".typo-rel")
	if typos.Length() == 0 {
		return nil
	}
	result := [][]string{}
	typos.Each(func(_ int, s *goquery.Selection) {
		word := strings.TrimSpace(s.Find("a").Text())
		s.Children().Remove()
		mean := strings.TrimSpace(s.Text())
		result = append(result, []string{word, mean})
	})
	return result
}
