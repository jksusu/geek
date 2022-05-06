package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "查看版本号",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 0.01")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
