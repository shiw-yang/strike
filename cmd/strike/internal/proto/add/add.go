package add

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// AddCmd represents the add command.
var AddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a proto API template",
	Long:    "Add a proto API template",
	Example: "strike proto add helloworld/v1/foo.proto",
	Run:     run,
}

func run(cmd *cobra.Command, args []string) {
	// strike proto add helloworld/v1/helloworld.proto
	if len(args) > 1 {
		fmt.Println("Too many args entered")
		return
	}
	input := args[0]
	index := strings.LastIndex(input, "/")
	if index == -1 {
		fmt.Println("The proto path needs to be hierarchical.")
		return
	}
	path := input[:index]
	filename := input[index+1:]
	filename, ok := checkProtoFileFromat(filename)
	if !ok {
		fmt.Println("filename is illegal.")
		return
	}
	pkgName := strings.ReplaceAll(path, "/", ".")
	p := &Proto{
		Name:      filename,
		Path:      path,
		Package:   pkgName,
		GoPackage: goPackage(path),
		Service:   serviceName(filename),
	}
	if err := p.Generate(); err != nil {
		fmt.Println(err)
		return
		
	}
}

// checkProtoFileFromat check the filename like xxx.proto
func checkProtoFileFromat(filename string) (file string, ok bool) {
	count := strings.Count(filename, ".")
	if count > 1 {
		return "", false
	}
	index := strings.LastIndex(filename, ".")
	if index == -1 {
		return filename + ".proto", true
	}
	if filename[index+1:] != "proto" {
		return "", false
	}
	return filename, true
}

func modName() string {
	modBytes, err := os.ReadFile("go.mod")
	if err != nil {
		if modBytes, err = os.ReadFile("../go.mod"); err != nil {
			return ""
		}
	}
	return modfile.ModulePath(modBytes)
}

func goPackage(path string) string {
	s := strings.Split(path, "/")
	return modName() + "/" + path + ";" + s[len(s)-1]
}

func serviceName(name string) string {
	return toUpperCamelCase(strings.Split(name, ".")[0])
}

func toUpperCamelCase(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = cases.Title(language.Und, cases.NoLower).String(s)
	return strings.ReplaceAll(s, " ", "")
}
