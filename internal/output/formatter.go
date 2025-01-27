package output

import "github.com/seankim658/depend-ls/internal/core"

type Formatter interface {
	Format([]*core.Dependency) (string, error)
}
