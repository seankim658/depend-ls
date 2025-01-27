package main

import (
	lang "depend-ls/languages"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var filePath = flag.String("file", "", "Path to the file to analyze")
	flag.Parse()

	if *filePath == "" {
		log.Fatal("Please provide a file path")
	}

	content, err := os.ReadFile(*filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	ext := filepath.Ext(*filePath)
	var language lang.Language
	switch ext {
	case ".py":
		language = lang.NewPythonLanguage()
	default:
		log.Fatalf("Unsupported file type: %s", ext)
	}


	parser := NewParser(language)
	analyzer := NewAnalyzer(parser)

	deps, err := analyzer.AnalyzeFile(content)
	if err != nil {
		log.Fatalf("Error analyzing file: %v", err)
	}

	writeDependencies(deps)
}

func writeDependencies(deps []*Dependency) {
	for _, dep := range deps {
		fmt.Printf("## %s: %s (line %d)\n", dep.Type, dep.Name, dep.Line)

		if len(dep.Calls) > 0 {
			fmt.Println("Calls:")
			for name, refs := range dep.Calls {
				for _, ref := range refs {
					fmt.Printf("- %s\n  - Used at line %d\n", name, ref.Line)
				}
			}
		}
		fmt.Println()
	}
}
