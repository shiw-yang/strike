package project

import "github.com/spf13/cobra"

// CmdCommand represents the new command.
var CmdCommand = &cobra.Command{
	Use:   "New",
	Short: "Create a service template",
	Long:  "Create a service project the repository template. Example: strike new helloworld",
	Run:   run,
}

// run is the impl of CmdCommand.Run
func run(cmd *cobra.Command, args []string) {
	//todo
}
