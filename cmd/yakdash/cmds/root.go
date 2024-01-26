package cmds

import (
	"log"

	"github.com/spf13/cobra"
)

func Execute() error {
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use: "yakdash",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hello from yakdash!")
	},
}
