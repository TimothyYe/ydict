package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

const ()

func query(word string) {
	var url string
	isChinese := IsChinese(word)

	if isChinese {
		url = "http://dict.youdao.com/w/eng/%s"
	} else {
		url = "http://dict.youdao.com/w/%s"
	}

	doc, err := goquery.NewDocument(fmt.Sprintf(url, word))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if isChinese {
		// Find the result
		fmt.Println()
		doc.Find(".trans-container > ul > p > span.contentTitle").Each(func(i int, s *goquery.Selection) {
			color.Yellow("    %s", s.Find(".search-js").Text())
		})
	} else {
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

func main() {
	if len(os.Args) == 1 {
		displayUsage()
		os.Exit(0)
	}

	words := strings.Join(os.Args[1:], " ")
	query(words)
}
