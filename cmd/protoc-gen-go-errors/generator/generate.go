package generator

// go:embed template.go.tpl
var tpl string

// Exc todo
func Exc() string {
	return tpl
}
