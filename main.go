package main

//import "fmt"

type ReportRow struct {
	Name        string
	Version     string
	Link        string
	Description string
	License     string
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












// cmd := exec.Command("go", "list", "-deps", "-json")
	// out, err := cmd.Output() 
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Output of 'go list -json' is a stream of JSON objects, rather than a single
	// valid JSON object
	// dec := json.NewDecoder(strings.NewReader(string(out)))


	// for {
	// 	var goListLit GoListLiteral
	// 	if err := dec.Decode(&goListLit); err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		log.Fatal(err)
	// 	}
	
		// We shouldn't care about std packages for purposes of a report
		// (among other possible reasons, because that code originates
		// in the system Go RPM, which would be reported separately from
		// the Go binary components being reported here)
		// Not currently sure if checking value of 'Standard' will catch
		// everything we want to exclude as a 'system' package

		// Also, we only want to provide information on Go modules
		// TODO: check if it's duplicative to check for both 'Standard == false' and
		// 'Module != nil'
		
	// 	if !goListLit.Standard && goListLit.Module != nil {
	// 		modSlice = append(modSlice, *goListLit.Module)
	// 	}

	// }

// 	PopulateReport(&modSlice, &report)
	
// }



