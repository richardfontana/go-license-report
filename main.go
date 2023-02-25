package main

import (
	"fmt"
)

type ReportRow struct {
	Name        string
	Version     string
	Url         string
	License     string
	Description string
}

type Module struct {
	Path string
	Version string
	Replace *Module
	Main bool
}

type FlatMod struct {
	Path string
	Version string
	ReplacePath string
	ReplaceVersion string
	Main bool
}

func ModuleToFlatMod(module Module) FlatMod {
	var f FlatMod
	f.Path = module.Path
	f.Version = module.Version
	if module.Replace != nil {
		f.ReplacePath = module.Replace.Path
		f.ReplaceVersion = module.Replace.Version
	}
	f.Main = module.Main
	return f
}

func main() {
	//	var report []ReportRow

	fmap := make(map[FlatMod]bool)
	
	GetModulesUsed(&fmap)

	for r := range fmap {	
		fmt.Println(r)
	}


	
	//	PopulateReport(&mods, &report)
	//	CacheResults(report)
	//	WriteToCSV(report)
}	
