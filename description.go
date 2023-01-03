package main

import (
	_	"embed"
	"strings"
	"github.com/neurosnap/sentences"
	"github.com/PuerkitoBio/goquery"
)

// data/english.json file copied from github.com/neurosnap/sentences

//go:embed data/english.json
var trainData []byte

func Description(doc *goquery.Document) string {

	readme := doc.Find(".Overview-readmeContent").
		ChildrenFiltered("p").
		FilterFunction(func(i int, s *goquery.Selection) bool {
			return strings.TrimSpace(s.Text()) != ""
		}).First()

	if readme.Text() == "" {
		return ""
	} else {
		return FirstTokenizedSentence(readme.Text())
	}
}

func FirstTokenizedSentence(text string) string {

	// load training data
	training, _ := sentences.LoadTraining(trainData)

	// create default sentence tokenizer
	tokenizer := sentences.NewSentenceTokenizer(training)
	sentences := tokenizer.Tokenize(text)

	return sentences[0].Text
}

