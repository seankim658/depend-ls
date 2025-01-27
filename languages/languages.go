package languages

import sitter "github.com/smacker/go-tree-sitter"

// Represents a programming language that can be analyzed.
type Language interface {
	// Returns the tree-sitter language implementation.
	GetTreeSitterLanguage() *sitter.Language
	// Return language-specific tree-sitter queries.
	GetQueries() LanguageQueries
}

// The tree-sitter queries needed to analyze a language.
type LanguageQueries struct {
	FunctionDefinition string
	FunctionCalls      string
	TypeReferences     string
	ConstantReferences string
}
