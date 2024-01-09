package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/Twacqwq/gpt-terminal/bootstrap"
	"github.com/Twacqwq/gpt-terminal/internal/command"
	"github.com/Twacqwq/gpt-terminal/internal/driver"
	"github.com/Twacqwq/gpt-terminal/internal/typs"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "gptchat",
	Short: "A terminal chat GPT that integrates major models",
	Long:  "",
	Run:   run,
}

func init() {
	rootCmd.AddCommand(
		cmdList,
		cmdConfig,
	)
}

func Execute() error {
	return rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	defer command.Clean()
	if err := driver.Viper().ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, color.RedString("💥 Please run `gptchat init-config` to initialize the configuration file first."))
		println(err.Error())
		return
	}
	fmt.Fprintln(os.Stderr, color.CyanString("⚡️ A terminal chat GPT that integrates major models"))
	fmt.Fprintln(os.Stderr, color.GreenString("Config is in `%s`", driver.Viper().ConfigFileUsed()))

	chExit := make(chan struct{}, 1)

	// listen terminal stdin
	for {
		select {
		case <-chExit:
			close(chExit)
			return
		default:
			line, err := command.Rl.Readline()
			if err != nil {
				break
			}
			listen(cmd.Context(), chExit, line)
		}
	}
}

func listen(ctx context.Context, chExit chan struct{}, line string) {
	switch line {
	case typs.Help:
		color.YellowString("🔨 Comming soon...")
	case typs.MiniMax:
		terminal, err := bootstrap.CreateTerminal(ctx, line)
		if err != nil {
			color.RedString("%v", err.Error())
			return
		}
		terminal.Prompt(ctx)
		for {
			line, err = command.Rl.Readline()
			if err != nil {
				break
			}
			if line == typs.Back {
				command.Rl.SetPrompt("\033[31mhost»\033[0m ")
				return
			}
			err := terminal.Chat(ctx, line)
			if err != nil {
				break
			}
			println()
		}
	case typs.Exit:
		fmt.Fprintln(command.Rl.Stderr(), color.GreenString("Bye~ 🙈"))
		command.Rl.CaptureExitSignal()
		chExit <- struct{}{}
		return
	default:
		fmt.Fprintln(os.Stderr, color.RedString("💥 Command not found."))
	}
}
