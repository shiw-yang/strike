package project

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// NewCmd represents the new command.
var NewCmd = &cobra.Command{
	Use:     "new",
	Short:   "Create a service template",
	Long:    "Create a service project the repository template.",
	Example: "strike new helloworld",
	Run:     run,
}
var (
	repoURL string
	timeout string
)

func init() {
	if repoURL = os.Getenv("STRIKE_LAYOUT_REPO"); repoURL == "" {
		repoURL = "https://github.com/shiw-yang/strike_layout_template.git"
	}
	timeout = "60s"
	NewCmd.Flags().StringVarP(&repoURL, "repo-url", "r", repoURL, "layout repo")
	NewCmd.Flags().StringVarP(&timeout, "time-out", "t", timeout, "time out")
}

// run is the impl of CmdCommand.Run
func run(cmd *cobra.Command, args []string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// time out setting
	t, err := time.ParseDuration(timeout)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	// get project name
	projectName := ""
	if len(args) == 0 {
		prompt := &survey.Input{
			Message: "What is project name?",
			Default: "helloworld",
			Help:    "Create project name.",
		}
		err = survey.AskOne(prompt, &projectName)
		if err != nil {
			return
		}
	} else {
		projectName = args[0]
	}
	projectName, workingDir := processProjectDir(projectName, dir)
	project := &Project{
		Name: projectName,
		Path: projectName,
	}

	// download template
	done := make(chan error, 1)
	go func() {
		done <- project.New(ctx, workingDir, repoURL)
	}()

	// check download status
	select {
	case <-ctx.Done():
		if err := ctx.Err(); err != nil {
			fmt.Fprint(os.Stderr, "\033[31mERROR: project creation timed out\033[m\n")
			return
		}
		fmt.Fprintf(os.Stderr, "\033[31mERROR: failed to create project(%s)\033[m\n", projectName)
	case err := <-done:
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[31mERROR: failed to create project(%s)\033[m\n", projectName)
			return
		}
	}

}

// processProjectDir process project name and dir, when name is absolute path, dir is ignored.
func processProjectDir(name, dir string) (projectName string, workingDir string) {
	_name := name
	_dir := dir
	// check name path is relative or absolute
	if !filepath.IsAbs(name) {
		absPath, err := filepath.Abs(name)
		if err != nil {
			return _name, _dir
		}
		_name = absPath
	}

	return filepath.Base(_name), filepath.Dir(_name)
}
