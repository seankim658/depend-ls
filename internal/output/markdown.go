package output

import (
	"fmt"
	"github.com/seankim658/depend-ls/internal/core"
  "strings"
)

type MarkdownFormatter struct{}

func (m *MarkdownFormatter) Format(deps []*core.Dependency) (string, error) {
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
