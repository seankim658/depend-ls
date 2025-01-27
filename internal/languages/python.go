package languages

import (
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/python"
)

// Implements the Language interface for Python.
type PythonLanguage struct{}

func NewPythonLanguage() *PythonLanguage {
	return &PythonLanguage{}
}

func (p *PythonLanguage) GetTreeSitterLanguage() *sitter.Language {
	return python.GetLanguage()
}

func (p *PythonLanguage) GetQueries() LanguageQueries {
	return LanguageQueries{
		FunctionDefinition: `
            (function_definition
                name: (identifier) @func.name
                body: (block) @func.body)
        `,
		FunctionCalls: `
            (call
                function: (identifier) @call.name)
        `,
		TypeReferences: `
            (type [(identifier) (attribute)] @type.name)
        `,
		ConstantReferences: `
            (identifier) @const.name
            (#match? @const.name "^[A-Z][A-Z_]*$")
        `,
	}
}
