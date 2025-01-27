package main

import (
	"context"
	lang "depend-ls/languages"
	"fmt"
	sitter "github.com/smacker/go-tree-sitter"
)

// Handles the tree-sitter parsing of source code.
type Parser struct {
	parser   *sitter.Parser
	language lang.Language
}

// Creates a new Parser instance for the specified language.
func NewParser(lang lang.Language) *Parser {
	parser := sitter.NewParser()
	parser.SetLanguage(lang.GetTreeSitterLanguage())

	return &Parser{parser: parser, language: lang}
}

// Entry point for analyzing a file's content.
func (p *Parser) ParseFile(content []byte) ([]*Dependency, error) {
	tree, err := p.parser.ParseCtx(context.Background(), nil, content)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %w", err)
	}
	defer tree.Close()

	queries := p.language.GetQueries()

	//// Function Calls

	// First find all functions and their corresponding body nodes
	functions, err := p.parseFunctions(content, tree.RootNode(), queries.FunctionDefinition)
	if err != nil {
		return nil, fmt.Errorf("error parsing functions: %w", err)
	}
	// Then find all function calls within each function
	for _, fn := range functions {
		calls, err := p.findFunctionCalls(content, fn.bodyNode, queries.FunctionCalls)
		if err != nil {
			return nil, fmt.Errorf("error finding function calls: %w", err)
		}
		fn.Calls = calls
	}

	return functions, nil
}

// Analyzes the content and returns the found dependencies.
func (p *Parser) parseFunctions(content []byte, root *sitter.Node, query string) ([]*Dependency, error) {
	q, err := sitter.NewQuery([]byte(query), p.language.GetTreeSitterLanguage())
	if err != nil {
		return nil, fmt.Errorf("function query error: %w", err)
	}
	defer q.Close()

	var functions []*Dependency
	qc := sitter.NewQueryCursor()
	qc.Exec(q, root)

	for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}

		var fnName string
		var bodyNode *sitter.Node

		for _, c := range m.Captures {
			switch c.Index {
			case 0:
				fnName = c.Node.Content(content)
			case 1:
				bodyNode = c.Node
			}
		}

		if fnName != "" && bodyNode != nil {
			fn := &Dependency{
				Name:      fnName,
				Type:      "function",
				Line:      bodyNode.StartPoint().Row + 1,
				Calls:     make(map[string][]Reference),
				UsesTypes: make(map[string][]Reference),
				Constants: make(map[string][]Reference),
				bodyNode:  bodyNode,
			}
			functions = append(functions, fn)
		}
	}

	return functions, nil
}

func (p *Parser) findFunctionCalls(content []byte, node *sitter.Node, query string) (map[string][]Reference, error) {
	q, err := sitter.NewQuery([]byte(query), p.language.GetTreeSitterLanguage())
	if err != nil {
		return nil, fmt.Errorf("call query error: %w", err)
	}
	defer q.Close()

	calls := make(map[string][]Reference)
	qc := sitter.NewQueryCursor()
	qc.Exec(q, node)

	for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}

		for _, c := range m.Captures {
			// For function calls, we want the identifier nodes
			if c.Node.Type() == "identifier" {
				name := c.Node.Content(content)
				ref := Reference{
					Line:   c.Node.StartPoint().Row + 1,
					Column: c.Node.StartPoint().Column + 1,
				}
				calls[name] = append(calls[name], ref)
			}
		}
	}

	return calls, nil
}
