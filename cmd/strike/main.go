package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "strike",
	Short:   "Strike: An simple toolkit for Go microservices.",
	Long:    `Strike: An simple toolkit for Go microservices.`,
	Version: release,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
