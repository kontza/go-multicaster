package cmd

import (
	"github.com/spf13/cobra"
)

func rootRunner(cmd *cobra.Command, args []string) {
	cmd.Usage()
}
