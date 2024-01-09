package command

import (
	"sync"

	"github.com/chzyer/readline"
)

var (
	Rl   *readline.Instance
	once sync.Once
)

func init() {
	getReadLine()
}

func Clean() {
	Rl.Close()
}

func getReadLine() {
	once.Do(func() {
		cli, err := readline.NewEx(&readline.Config{
			Prompt:                 "\033[31mgptchat»\033[0m ",
			DisableAutoSaveHistory: true,
			InterruptPrompt:        "^C",
			EOFPrompt:              "exit",
		})
		if err != nil {
			panic(err)
		}
		Rl = cli
	})
}
