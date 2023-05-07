package upgrade

import (
	"fmt"

	"github.com/shiw-yang/strike/cmd/strike/internal/base"
	"github.com/spf13/cobra"
)

// UpgradeCmd represents the upgrade command.
var UpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the strike tools",
	Long:  "Upgrade the strike tools. Example: strike upgrade",
	Run:   Run,
}

// Run upgrade the strike tools.
func Run(_ *cobra.Command, _ []string) {
	err := base.GoInstall(
		"github.com/shiw-yang/strike/cmd/strike@latest",
		"github.com/shiw-yang/strike/cmd/protoc-gen-go-errors@latest",
		"github.com/shiw-yang/strike/cmd/protoc-gen-gin@latest",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
		"github.com/google/gnostic/cmd/protoc-gen-openapi@latest",
		"github.com/favadi/protoc-go-inject-tag@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}
