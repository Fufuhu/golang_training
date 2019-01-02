package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "(Short) Print the version number of this application",
	Long:  "(Long) Print the version number of this application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This Application version is v0.0.1")
	},
}
