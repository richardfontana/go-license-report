package main

import (
	"encoding/json"
	//	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

func GetModulesUsed(fmap *map[FlatMod]bool) {

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

	cmd := exec.Command("go", "list", "-e", "-deps", "-json", "...")
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
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
		// We only report on non-std modules and ignore golang.org/x/ modules
		if !g.Standard && g.Module != nil &&
			!strings.HasPrefix(g.Module.Path, "golang.org/x/") {
			f := ModuleToFlatMod(*g.Module)
			(*fmap)[f] = true
		}
	}
}


