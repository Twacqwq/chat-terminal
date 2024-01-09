package minimax

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/Twacqwq/go-minimax"
	"github.com/Twacqwq/gpt-terminal/internal/command"
	"github.com/Twacqwq/gpt-terminal/internal/driver"
)

type MiniMax struct {
	client *minimax.Client
}

func NewMiniMax() *MiniMax {
	return &MiniMax{}
}

func (m *MiniMax) Init(ctx context.Context) error {
	v := driver.Viper()
	accessToken := v.GetString("models.minimax.accesstoken")
	groupId := v.GetString("models.minimax.groupid")
	cli, err := miniMaxClient(v.GetString("models.minimax.accesstoken"), v.GetString("models.minimax.groupid"))
	if accessToken == "" || groupId == "" {
		return fmt.Errorf("invalid minimax config")
	}
	m.client = cli

	return err
}

func (m *MiniMax) Prompt(ctx context.Context) {
	command.Rl.SetPrompt("\033[31mgptchat->minimax»\033[0m ")
}

func (m *MiniMax) Chat(ctx context.Context, message string) error {
	resp, err := m.client.CreateCompletionStream(ctx, &minimax.ChatCompletionRequest{
		Model: minimax.Abab5Dot5,
		Messages: []minimax.Message{
			{
				SenderType: minimax.ChatMessageRoleUser,
				SenderName: driver.Viper().GetString("models.minimax.username"),
				Text:       message,
			},
		},
	})
	if err != nil {
		return err
	}
	defer resp.Close()
	for {
		res, err := resp.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		if res.Choices[0].FinishReason == "" {
			fmt.Fprint(command.Rl.Stderr(), res.Choices[0].Messages[0].Text)
		}
	}

	return nil
}

var (
	_ driver.Driver      = (*MiniMax)(nil)
	_ driver.GPTTerminal = (*MiniMax)(nil)
)
