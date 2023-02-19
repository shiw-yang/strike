package proto

import (
	"github.com/spf13/cobra"
	"github.com/shiw-yang/strike/cmd/strike/internal/proto/add"
)

// ProtoCmd represents the proto command.
var ProtoCmd = &cobra.Command{
	Use:   "proto",
	Short: "Generate the proto files",
	Long:  "Generate the proto files.",
}

func init() {
	ProtoCmd.AddCommand(add.AddCmd)
}
