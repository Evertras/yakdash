package cmds

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/yakdash/pkg/yakdash"
	"github.com/spf13/cobra"
)

func Execute() error {
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use: "yakdash",
	Run: func(cmd *cobra.Command, args []string) {
		model := yakdash.New()

		p := tea.NewProgram(model)

		if _, err := p.Run(); err != nil {
			log.Fatal("Failed to run:", err)
		}
	},
}
