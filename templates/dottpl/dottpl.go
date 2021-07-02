package dottpl

import (
	"context"
	"embed"
	"fmt"
	"text/template"

	"github.com/xo/xo/templates"
	xo "github.com/xo/xo/types"
)

func init() {
	templates.Register("dot", &templates.TemplateSet{
		Files:   Files,
		FileExt: ".xo.dot",
		Flags: []templates.Flag{
			{
				ContextKey:  DefaultsKey,
				Desc:        "default statements, default: node [shape=none, margin=0]",
				PlaceHolder: `""`,
				Default:     "node [shape=none, margin=0]",
				Value:       []string{},
			},
			{
				ContextKey: BoldKey,
				Desc:       "bold header row",
				Default:    "false",
				Value:      false,
			},
			{
				ContextKey: ColorKey,
				Desc:       "header color",
				Default:    "lightblue",
				Value:      "",
			},
			{
				ContextKey:  RowKey,
				Desc:        "row value template, default:  {{ .Name }}: {{ .Datatype.Type }}",
				Default:     "{{ .Name }}: {{ .Datatype.Type }}",
				PlaceHolder: `""`,
				Value:       "",
			},
			{
				ContextKey: DirectionKey,
				Desc:       "enable edge directions",
				Default:    "true",
				Value:      true,
			},
		},
		Funcs: func(ctx context.Context) (template.FuncMap, error) {
			f, err := NewFuncs(ctx)
			if err != nil {
				return nil, err
			}
			return f.FuncMap(), nil
		},
		FileName: func(ctx context.Context, tpl *templates.Template) string {
			return tpl.Name
		},
		Process: func(ctx context.Context, _ bool, set *templates.TemplateSet, v *xo.XO) error {
			if len(v.Schemas) == 0 {
				return fmt.Errorf("dot output only works with schema mode")
			}
			for _, schema := range v.Schemas {
				if err := set.Emit(ctx, &templates.Template{
					Name:     "xo",
					Template: "xo",
					Data:     schema,
				}); err != nil {
					return err
				}
			}
			return nil
		},
		Order: []string{"xo"},
	})
}

// Context keys.
const (
	DefaultsKey  xo.ContextKey = "defaults"
	DirectionKey xo.ContextKey = "direction"
	BoldKey      xo.ContextKey = "bold"
	RowKey       xo.ContextKey = "row"
	ColorKey     xo.ContextKey = "color"
)

// Defaults returns default values from the context.
func Defaults(ctx context.Context) []string {
	s, _ := ctx.Value(DefaultsKey).([]string)
	return s
}

// Bold returns bold from the context.
func Bold(ctx context.Context) bool {
	b, _ := ctx.Value(BoldKey).(bool)
	return b
}

// Color returns color from the context.
func Color(ctx context.Context) string {
	s, _ := ctx.Value(ColorKey).(string)
	return s
}

// Row returns the row template from the context.
func Row(ctx context.Context) string {
	s, _ := ctx.Value(RowKey).(string)
	return s
}

// Direction returns edge direction from the context.
func Direction(ctx context.Context) bool {
	b, _ := ctx.Value(DirectionKey).(bool)
	return b
}

// Files are the embedded dot templates.
//
//go:embed *.tpl
var Files embed.FS