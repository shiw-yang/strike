package generator

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed template.go.tpl
var errorsTpl string

type errorInfo struct {
	Name       string
	Value      string
	HTTPCode   int
	CamelValue string
	Comment    string
	HasComment bool
}

type errorWrapper struct {
	Errors []*errorInfo
}

func (e *errorWrapper) execute() string {
	buf := &bytes.Buffer{}
	tmpl, err := template.New("errors").Parse(errorsTpl)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, e); err != nil {
		panic(err)
	}
	return buf.String()
}
