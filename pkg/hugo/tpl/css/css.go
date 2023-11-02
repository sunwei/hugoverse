package css

import (
	"context"

	"github.com/spf13/cast"
	"github.com/sunwei/hugoverse/pkg/hugo/common/types/css"
	"github.com/sunwei/hugoverse/pkg/hugo/deps"
	"github.com/sunwei/hugoverse/pkg/hugo/tpl/internal"
)

const name = "css"

// Namespace provides template functions for the "css" namespace.
type Namespace struct {
}

// Quoted returns a string that needs to be quoted in CSS.
func (ns *Namespace) Quoted(v any) css.QuotedString {
	s := cast.ToString(v)
	return css.QuotedString(s)
}

// Unquoted returns a string that does not need to be quoted in CSS.
func (ns *Namespace) Unquoted(v any) css.UnquotedString {
	s := cast.ToString(v)
	return css.UnquotedString(s)
}

func init() {
	f := func(d *deps.Deps) *internal.TemplateFuncsNamespace {
		ctx := &Namespace{}

		ns := &internal.TemplateFuncsNamespace{
			Name:    name,
			Context: func(cctx context.Context, args ...any) (any, error) { return ctx, nil },
		}

		return ns
	}

	internal.AddTemplateFuncsNamespace(f)
}
