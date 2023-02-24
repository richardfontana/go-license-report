package main

import (
	"fmt"
	"go/build"
	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
)


func ScrapeSecondaryData(row *ReportRow) {

	const PkgSite = "https://pkg.go.dev/"

	const Replaced = " < "

	// default to "UNKNOWN" to take care of the 404 case
	row.License = "UNKNOWN"
	if row.Url == "" {
		row.Url = "UNKNOWN"
	}
	row.Description = "UNKNOWN"

	var gen string

	wkgVer := row.Version
	
	replName, origName, nameReplaced := strings.Cut(row.Name, Replaced)
	replVer, origVer, verReplaced  := strings.Cut(row.Version, Replaced)


	gen = PkgSite + row.Name	
	
	if nameReplaced {
		if build.IsLocalImport(replName) {
			gen = PkgSite + origName
		} else {
			gen = PkgSite + replName
		}
	}

	if verReplaced {
	
		if replVer == "" {
			wkgVer = origVer
		} else {
			wkgVer = replVer
		}

	}
	
	gres, err := http.Get(gen)



	if err != nil {
		log.Fatal(err)
	}

	defer gres.Body.Close()
	
	gsc := gres.StatusCode

	var doc *goquery.Document

	if gsc == 200 {



		if wkgVer != "MAIN" {

			// try to get version-specific page
			
			ver := gen + "@" + wkgVer
			vres, err := http.Get(ver)
			if err != nil {
				log.Fatal(err)
			}
			defer vres.Body.Close()		
			vsc := vres.StatusCode

			if vsc == 200 {
		
				doc, err = goquery.NewDocumentFromReader(vres.Body)
			
			} else if vsc == 404 {

				doc, err = goquery.NewDocumentFromReader(gres.Body)	
			} else if vsc == 504 {
				// do nothing
								
			} else {
				fmt.Println("ver is ", ver)
				log.Fatalf("status code error: %d, %s", vsc, vres.Status)
			}
		
		} else { // row.Version == "MAIN" {
			doc, err = goquery.NewDocumentFromReader(gres.Body)
		}
		row.License = License(doc)
		if row.Url == "UNKNOWN" { // not a main module
			row.Url = Link(doc)
		}
		row.Description = Description(doc)
			
	} else if gsc == 404 {

		// for the replace case, try using the original path

		if nameReplaced {

			gen = PkgSite + origName
			ores, err := http.Get(gen)

			if err != nil {
				log.Fatal(err)
			}
			defer ores.Body.Close()		
			osc := ores.StatusCode

			if osc == 200 {
				doc, err = goquery.NewDocumentFromReader(ores.Body)

				row.License = License(doc)
				if row.Url == "UNKNOWN" { // not a main module
					row.Url = Link(doc)
				}
				row.Description = Description(doc)

			} else if osc == 404 || osc == 504 {
				// do nothing
			} else {
				log.Fatalf("status code error: %d, %s", osc, ores.Status)
			}
			
		} else {
			// use initial values
		}

	} else if gsc == 504 {
		// do nothing?
	} else {
		fmt.Println("gen is ", gen)
		log.Fatalf("status code error: %d %s", gsc, gres.Status)
	}
}
