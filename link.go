package main

import (
	//	"fmt"
	"github.com/PuerkitoBio/goquery"

)

func Link(doc *goquery.Document) string {
	
	val, _ := doc.Find(".UnitMeta-repo").ChildrenFiltered("a").First().Attr("href")
	return val
}
