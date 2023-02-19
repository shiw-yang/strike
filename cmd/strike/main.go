package main

import (
	"log"

	"github.com/shiw-yang/strike/cmd/strike/internal/project"
	"github.com/shiw-yang/strike/cmd/strike/internal/proto"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "strike",
	Short:   "Strike: An simple toolkit for Go microservices.",
	Long:    `Strike: An simple toolkit for Go microservices.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(project.NewCmd)
	rootCmd.AddCommand(proto.ProtoCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
