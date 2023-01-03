package main

import (
	"encoding/csv"
	"log"
	"os"
	"reflect"
)

func WriteToCSV(r []ReportRow) {

	t := reflect.TypeOf(ReportRow{})

	rowSize := t.NumField()
		
	headings := make([]string, rowSize)
	for i := range headings {
		headings[i] = t.Field(i).Name
	}

	repSize := len(r)
	
	s := make([][]string, repSize+1)

	for i :=0; i < repSize; i++ {
		s[i] = make([]string, rowSize)
		rval := reflect.ValueOf(&r[i]).Elem()
		for j := 0; j < rowSize; j++ {	
			s[i][j] = rval.Field(j).String()
		}
	}

	hslice := make([][]string, 1)

	hslice[0] = make([]string, rowSize)
	
	hslice[0] = headings

	for i := 0; i < repSize; i++ {
		hslice = append (hslice, s[i])
	}

	file, err := os.Create("report.csv")
	if err != nil {
	 	log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	writer.Comma = '\t'
	
	defer writer.Flush()

	writer.WriteAll(hslice)

}
