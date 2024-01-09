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

	if err := v.WriteConfig(); err != nil {
		panic(err)
	}

	fmt.Fprintln(os.Stderr, color.GreenString("🚀 Config init done."))
}
