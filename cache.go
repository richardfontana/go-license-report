package main

import (
	"encoding/json"
	"log"
)


func CacheResults(r []ReportRow) {
	
	_, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	
}
