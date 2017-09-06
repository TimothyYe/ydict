package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func query(word string) {
	doc, err := goquery.NewDocument(fmt.Sprintf("http://dict.youdao.com/w/%s", word))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Find the result
	result := doc.Find("div#phrsListTab > div.trans-container > ul").Text()
	fmt.Println(result)
}

func main() {
	if len(os.Args) == 1 {
		displayUsage()
		os.Exit(0)
	}

	word := os.Args[1]
	query(word)
}
