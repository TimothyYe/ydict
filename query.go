package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	proxier "golang.org/x/net/proxy"
)

func query(words []string, playVoice, isMulti bool) {
	var url string
	var doc *goquery.Document
	queryString := strings.Join(words, " ")
	//voiceString := strings.Join(words, "+")

	isChinese := isChinese(queryString)

	if isChinese {
		url = "http://dict.youdao.com/w/eng/%s"
	} else {
		url = "http://dict.youdao.com/w/%s"
	}

	//Init spinner
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = "Querying... "
	s.Color("green")
	s.Start()

	//Check proxy
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

		resp, err := client.Get(fmt.Sprintf(url, queryString))

		if err != nil {
			color.Red("Query failed with err: %s", err.Error())
			os.Exit(1)
		}

		doc, _ = goquery.NewDocumentFromResponse(resp)
	} else {
		var err error
		doc, err = goquery.NewDocument(fmt.Sprintf(url, queryString))
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	s.Stop()

	if isChinese {
		// Find the result
		fmt.Println()
		doc.Find(".trans-container > ul > p > span.contentTitle").Each(func(i int, s *goquery.Selection) {
			color.Green("    %s", s.Find(".search-js").Text())
		})
	} else {
		// Find the pronounce
		if !isMulti {
			color.Green("\r\n    %s", getPronounce(doc))
		}

		// Find the result
		result := doc.Find("div#phrsListTab > div.trans-container > ul").Text()
		color.Green(result)
	}

	// Show examples
	sentences := getSentences(doc, isChinese)
	if len(sentences) > 0 {
		fmt.Println()
		for _, sentence := range sentences {
			color.Green("    %s", sentence[0])
			color.Magenta("    %s", sentence[1])
		}
		fmt.Println()
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

func getSentences(doc *goquery.Document, isChinese bool) [][]string {
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
					addSpace := (ii == 1 && isChinese) || (ii == 0 && !isChinese)
					if addSpace && iii != 0 && text != "." {
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
