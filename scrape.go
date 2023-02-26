package main

import (
	//	"fmt"
	"go/build"
	//	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
)
func ScrapeSecData(row *ReportRow) {

	row.License = "UNKNOWN"
	row.Url = "UNKNOWN"
	row.Description = "UNKNOWN"

	const PkgSite = "https://pkg.go.dev/"
	const Replaced = " < "

	// Cut slices s around the first instance of sep, returning the
	// text before and after sep. The found result reports whether
	// sep appears in s. If sep does not appear in s, cut returns s,
	// "", false.

	// func CutSuffix(s, suffix string) (before string, found bool)
	// CutSuffix returns s without the provided ending suffix string and
	// reports whether it found the suffix. If s doesn't end with suffix,
        // CutSuffix returns s, false. If suffix is the empty string, CutSuffix
        // returns s, true.
	
	curPath, origPath, pathReplaced := strings.Cut(row.Name, Replaced)

	// In the Main case, this will result in:
	// curVer == "MAIN-UNKNOWN", origVer == "", Replaced == false, 
	curVer, origVer, verReplaced := strings.Cut(row.Version, Replaced)

	// These are probably the only cases:
	// 1. No replace in 'Name', no replace in 'Version'
	//      Main is a special case here
	//      
	// 2. No replace in 'Name', replace in 'Version'
	// 3. Replace in 'Name', replace in 'Version'
	// There shouldn't be a case where 'Name' is a replace but 'Version' is not a replace
	
	
	// In case 1 and case 2, first try to access https://pkg.go.dev/{curPath}@{curVer}
	// If that doesn't work, try https://pkg.go.dev/{curPath}
	// Special Main case: just try https://pkg.go.dev/{curPath}
	
	// In case 3, first check that replacePath is a local filepath (if not,
	// something is probably wrong)
	// There should be a replace version in this case (if not, something is
	// probably wrong)
	// We assume the replace version is not a valid version for the original path
	// Try to access https://pkg.go.dev/{origPath}@{origVer}
	// If that doesn't work, try https://pkg.go.dev/{origPath}
	
	var candidateUrl string
	var fallbackUrl string
	suffix := ""

	if !pathReplaced { // cases 1 and 2
		if curVer != "MAIN-UNKNOWN" {
			suffix = "@" + curVer
		}
		fallbackUrl = PkgSite + curPath
	} else { // case 3
		if !verReplaced || !build.IsLocalImport(curPath) {
			panic("Something went wrong")
		} else {
			suffix = "@" + origVer
			fallbackUrl = PkgSite + origPath
		}
	}	
	candidateUrl = fallbackUrl + suffix

	// In the Main case, candidateUrl == fallbackUrl
	res, _ := http.Get(candidateUrl)
	defer res.Body.Close()
	statusCode := res.StatusCode

	var doc *goquery.Document
	
	if statusCode != 200 {
		if curVer == "MAIN-UNKNOWN" {
			return
		} else {
			res, _ = http.Get(fallbackUrl)
			defer res.Body.Close()
			statusCode = res.StatusCode
			if statusCode != 200 {
				return
			}
		}
	}

	doc, _ = goquery.NewDocumentFromReader(res.Body)
	row.Url = Link(doc)
	row.License = License(doc)
	row.Description = Description(doc)
}			
