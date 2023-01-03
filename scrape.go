package main

import (
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)


func ScrapeSecondaryData(row *ReportRow) {

	const PkgSite = "https://pkg.go.dev/"

	gen := PkgSite + row.Name

	var res *http.Response
	var err error
	
	if row.Version == "MAIN" || row.Version == "" {

		res, err = http.Get(gen)
		if err != nil {
			log.Fatal(err)
		}
		
	} else {

		ver := gen + "@" + row.Version
		
		res, err = http.Get(ver)
		if err != nil {
			log.Fatal(err)
		}

		if res.StatusCode != 200 { // Try again with general page for pkg
			res, err = http.Get(gen)
			if err != nil {
				log.Fatal(err)
			}
		
			if res.StatusCode != 200 {
				log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			}
		}
	}

	if res != nil {

		defer res.Body.Close()
		
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		row.License = License(doc)
		if row.Link == "" {
			row.Link = Link(doc)
		}
		row.Description = Description(doc)
	}
}	
	




