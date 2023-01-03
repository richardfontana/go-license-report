package main

import (
	"encoding/json"
	"io"
	"log"
	"os/exec"
	"strings"
)

func GetModulesUsed(m *[]Module) {

	type GoListLit struct {
		Dir        string
		ImportPath string
		Name       string
		Goroot     bool
		Standard   bool
		DepOnly    bool
		Root       string
		Module     *Module
	}

	cmd := exec.Command("go", "list", "-deps", "-json")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// 'go list -json' provides a stream of JSON objects,
	// rather than a single valid JSON object
	
	dec := json.NewDecoder(strings.NewReader(string(out)))

	for {
		var g GoListLit
		
		if err := dec.Decode(&g); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// We only report on non-std modules
		if !g.Standard && g.Module != nil {
			// We also ignore golang.org/x/ modules  
			if !strings.HasPrefix(g.Module.Path, "golang.org/x/") {
				*m = append(*m, *g.Module)
			}
		}
	}
}
