package main

import (
	"fmt"
	"log"
	"os"

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
			color.Blue("    %s", s.Find(".search-js").Text())
		})
		//result := doc.Find("div#phrsListTab > div.trans-container > ul.wordGroup").Text()
		//fmt.Println(result)
	} else {
		// Find the result
		result := doc.Find("div#phrsListTab > div.trans-container > ul").Text()
		color.Blue(result)
	}
}

func main() {
	if len(os.Args) == 1 {
		displayUsage()
		os.Exit(0)
	}

	word := os.Args[1]
	query(word)
}
