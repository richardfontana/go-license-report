package main

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func License(doc *goquery.Document) string {
	
	license := doc.Find("[data-test-id|='UnitHeader-license']").First().Text()

	// pkgsite reports some cases of conjunctive licensing using a pair of comma-separated
	// SPDX identifiers; let's put this in SPDX expression form
	return (strings.ReplaceAll(license, ",", " AND"))
	
}



