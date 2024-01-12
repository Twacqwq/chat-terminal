package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Twacqwq/gpt-terminal/internal/driver"
	"github.com/Twacqwq/gpt-terminal/internal/utils"
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

	if err := utils.Mkdir(filepath.Dir(v.ConfigFileUsed())); err != nil {
		panic(err)
	}

	if err := v.WriteConfig(); err != nil {
		panic(err)
	}

	fmt.Fprintln(os.Stderr, color.GreenString("🚀 Config init done."))
	fmt.Fprintln(os.Stderr, color.GreenString("use `cd $HOME/.gptchat && vim config`"))
}
