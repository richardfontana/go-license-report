package main

import (
	//	"fmt"
	"sort"
	"strings"
)	

func PopulateReport(mods *[]Module, rep *[]ReportRow) {

	for _, m := range *mods {

		var r ReportRow

		r.Name = m.Path
		r.Version = m.Version
		
		if m.Main {
			if m.Version == "" {
				r.Version = "MAIN"
			}
			r.Url = strings.Join([]string{"https://", r.Name}, "")
		}

		// Any non-main module should have a non-empty version (?)
		replace := m.Replace


		// if there is a replace directive, but m.Replace.Path is a relative
		// local filepath, and (probably in all such cases) there is no distinct
		// local version for the replace module,
		// we need to preserve the original information for the pkgsite scraping
		// and possibly to facilitate post-generation report cleanup


		if replace != nil {
			
			if r.Name != replace.Path {
				r.Name = replace.Path + " < " + r.Name
			}

			if r.Version != replace.Version {		
				r.Version = replace.Version + " < " + r.Version
			}

		}

		ScrapeSecondaryData(&r)
			
		*rep = append(*rep, r)
	}

	*rep = RemoveDuplicates(rep)
	
}

func RemoveDuplicates(r *[]ReportRow) []ReportRow {
 
 	sort.Slice(*r, func(i, j int) bool {
 		return (*r)[i].Name < (*r)[j].Name
	})

 	u := 0
	
 	for i := 1; i < len(*r); i++ {
 		if (*r)[u].Name != (*r)[i].Name ||
 			(*r)[u].Version != (*r)[i].Version {
 			u++
 			(*r)[u] = (*r)[i]
 		}
 	}

 	return (*r)[:u+1]
}

