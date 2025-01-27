package output

import (
	"fmt"
	"github.com/seankim658/depend-ls/internal/core"
	"strings"
)

type MarkdownFormatter struct{}

func NewMarkdownFormatter() *MarkdownFormatter {
	return &MarkdownFormatter{}
}

func (m *MarkdownFormatter) Format(deps []*core.Dependency) (string, error) {
	var sb strings.Builder

	for _, dep := range deps {
		fmt.Fprintf(&sb, "## %s: `%s` (line %d)\n", strings.ToTitle(dep.Type), dep.Name, dep.Line)

		if len(dep.Calls) > 0 {
			fmt.Fprintln(&sb, "Calls:")
      fmt.Fprintln(&sb)
			callIndex := 1
			for name, refs := range dep.Calls {
				for _, ref := range refs {
					fmt.Fprintf(&sb, "%d. `%s`\n  - Used at line %d\n", callIndex, name, ref.Line)
				}
				callIndex += 1
			}
		}
		fmt.Fprintln(&sb)
	}

	return sb.String(), nil
}
