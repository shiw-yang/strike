package project

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/shiw-yang/strike/cmd/strike/internal/base"
)

// Project is a project template.
type Project struct {
	Name string
	Path string
}

// New a project from remote repo.
func (p *Project) New(ctx context.Context, dir, layout string) error {
	// create project dir and download template
	to := filepath.Join(dir, p.Name)

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("🚫 %s already exists\n", p.Name)
		prompt := &survey.Confirm{
			Message: "📂 Do you want to override the folder ?",
			Help:    "Delete the existing folder and create the project.",
		}
		var override bool
		e := survey.AskOne(prompt, &override)
		if e != nil {
			return e
		}
		if !override {
			return err
		}
		os.RemoveAll(to)
	}
	// download the repo from remote
	fmt.Printf("🚀 Creating service %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)
	repo := base.NewRepo(layout)
	if err := repo.CopyTo(ctx, to, p.Path, []string{".git", ".github"}); err != nil {
		return err
	}
	// rename the cmd/server to cmd/projectName
	if renameErr := os.Rename(
		filepath.Join(to, "cmd", "server"),
		filepath.Join(to, "cmd", p.Name),
	); renameErr != nil {
		fmt.Println(renameErr)
		return renameErr
	}
	base.Tree(to, dir)

	fmt.Printf("\n🍺 Project creation succeeded %s\n", color.RedString(p.Name))
	fmt.Print("💻 Use the following command to start the project 👇:\n\n")

	return nil
}
