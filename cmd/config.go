package cmd

import (
	"fmt"
	"os"

	"github.com/Twacqwq/gpt-terminal/internal/driver"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cmdConfig = &cobra.Command{
	Use:   "init-config",
	Short: "init config",
	Run:   runInitConfig,
}

func runInitConfig(cmd *cobra.Command, args []string) {
	v := driver.Viper()
	v.Set("models", driver.NewConfig())

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	if err := v.WriteConfigAs(homeDir + "/.gptchat"); err != nil {
		panic(err)
	}

	fmt.Fprintln(os.Stderr, color.GreenString("🚀 Config init done."))
}
