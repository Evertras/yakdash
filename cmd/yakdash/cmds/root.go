package cmds

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/yakdash/pkg/yakdash"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func Execute() error {
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use: "yakdash",
	Run: func(cmd *cobra.Command, args []string) {
		model := yakdash.New(config.Layout)

		p := tea.NewProgram(model)

		if _, err := p.Run(); err != nil {
			log.Fatal("Failed to run:", err)
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	flags := rootCmd.Flags()

	flags.StringVarP(&cfgFile, "config", "c", "", "The config file to use")

	err := viper.BindPFlags(flags)
	if err != nil {
		log.Fatal("Failed to bind flags:", err)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// TODO: Default to XDG config home, and/or /etc somewhere
		log.Fatal("missing config file")
	}

	// Look for YAKDASH_*
	viper.SetEnvPrefix("yakdash")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Failed to parse config:", err)
	}
}
