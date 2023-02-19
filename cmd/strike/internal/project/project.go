package project

import (
	"fmt"
	"os"
	"path/filepath"

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
	// todo timeout
	// timeout string
)

func init() {
	if repoURL = os.Getenv("STRIKE_LAYOUT_REPO"); repoURL == "" {
		repoURL = "/home/shiwei/src/strike/layout"
	}
	// timeout = "60s"
	NewCmd.Flags().StringVarP(&repoURL, "repo-url", "r", repoURL, "layout repo")
	// NewCmd.Flags().StringVarP(&timeout, "time-out", "t", timeout, "time out")
}

// run is the impl of CmdCommand.Run
func run(cmd *cobra.Command, args []string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// todo time out
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
	wd := getProjectPlaceDir(projectName, dir)
	project := &Project{
		Name: filepath.Base(projectName),
		Path: projectName,
	}
	fmt.Printf("wd: %v\n", wd)
	fmt.Printf("project: %v\n", project)
}

func getProjectPlaceDir(name, dir string) string {
	panic("unimplemented")
}
