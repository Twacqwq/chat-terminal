package main

import (
	"github.com/Twacqwq/gpt-terminal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
