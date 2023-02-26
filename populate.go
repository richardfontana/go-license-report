package main

import (
	//	"fmt"
	//	"strings"
)	

func PopulateReport(fms *[]FlatMod, rep *[]ReportRow) {

	for _, f := range *fms {

		var r ReportRow

		// If there is a replace directive, but m.Replace.Path is a relative
		// local filepath, and (probably in all such cases) there is no distinct
		// local version for the replace module,
		// we need to preserve the original information for the pkgsite scraping
		// and possibly to facilitate post-generation report cleanup

		// In all the replace directives I've seen, the Path is present in the
		// Replace even if it's identical to the original Path

		// We (may) need to at least temporarily
		// remember somehow that there was a replace directive
		// for certain cases involving the 'Url' field

		if f.ReplacePath != "" && f.ReplacePath != f.Path {
			r.Name = f.ReplacePath + " < " + f.Path
		} else {
			r.Name = f.Path
		}

		if f.ReplaceVersion != "" && f.ReplaceVersion != f.Version {
			r.Version = f.ReplaceVersion + " < " + f.Version
		} else if f.Main && f.Version == "" {
			r.Version = "MAIN-UNKNOWN"
		} else {
			r.Version = f.Version
		}

		ScrapeSecData(&r)
		
		*rep = append(*rep, r)
	}

}

