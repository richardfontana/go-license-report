package main

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
	
func main() {
	var report []ReportRow
	var mods []Module
	
	GetModulesUsed(&mods)
	PopulateReport(&mods, &report)
	CacheResults(report)
	WriteToCSV(report)
}	
