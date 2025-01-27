package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/seankim658/depend-ls/internal/core"
	"github.com/seankim658/depend-ls/internal/languages"
	"github.com/seankim658/depend-ls/internal/output"
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
	var language languages.Language
	switch ext {
	case ".py":
		language = languages.NewPythonLanguage()
	default:
		log.Fatalf("Unsupported file type: %s", ext)
	}

	parser := core.NewParser(language)
	deps, err := parser.ParseFile(content)
	if err != nil {
		log.Fatalf("Error analyzing file: %v", err)
	}

	formatter := output.NewMarkdownFormatter()
	result, err := formatter.Format(deps)
	if err != nil {
		log.Fatalf("Error formatting output: %v", err)
	}

	fmt.Print(result)
}
