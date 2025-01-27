package main

import sitter "github.com/smacker/go-tree-sitter"

// Represents the specific location in code where an entity is used or defined.
type Reference struct {
	Line   uint32
	Column uint32
}

// Represents the type of dependency relationship.
type DependencyType string

const (
	FunctionCall DependencyType = "FUNCTION_CALL"
	TypeUsage    DependencyType = "TYPE_USAGE"
	ConstantRef  DependencyType = "CONSTANT_REG"
)

// Represents a code entity (function, class, etc.) and its dependencies.
type Dependency struct {
	// Name of the entity (e.g., function name)
	Name string
	// Type of the entity (e.g., "function", "class")
	Type string
	// Line where the entity is defined
	Line uint32
	// Function calls made by this entity
	Calls map[string][]Reference
	// Types used by this entity
	UsesTypes map[string][]Reference
	// Constants referenced by this entity
	Constants map[string][]Reference
	// Temporary field for processing
	bodyNode *sitter.Node
}

// Analyzer coordinates the parsing and analysis of source code.
type Analyzer struct {
	parser *Parser
}

// Creates a new Analyzer instance with the specified parser.
func NewAnalyzer(parser *Parser) *Analyzer {
	return &Analyzer{parser: parser}
}

// Processes a file's content and returns the dependencies.
func (a *Analyzer) AnalyzeFile(content []byte) ([]*Dependency, error) {
	return a.parser.ParseFile(content)
}
