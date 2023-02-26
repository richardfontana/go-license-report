package main

import (
	//	"fmt"
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

var fmSlice = make([]FlatMod, 0)

func FlatModToSlice(fmap *map[FlatMod]bool) {
	for flatmod := range *fmap {
		fmSlice = append(fmSlice, flatmod)
	}
}


func main() {
	var report []ReportRow

	fmap := make(map[FlatMod]bool)
	
	GetModulesUsed(&fmap)

	FlatModToSlice(&fmap)
	
	PopulateReport(&fmSlice, &report)
	WriteToCSV(report)
}	
