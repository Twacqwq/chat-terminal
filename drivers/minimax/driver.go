package minimax

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"

	"github.com/Twacqwq/go-minimax"
	"github.com/Twacqwq/gpt-terminal/internal/command"
	"github.com/Twacqwq/gpt-terminal/internal/driver"
	"github.com/Twacqwq/gpt-terminal/internal/storage"
)

type MiniMax struct {
	client *minimax.Client
	db     storage.SqliteRepo
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

	m.db = storage.OpenSqliteStorage(filepath.Join(filepath.Dir(v.ConfigFileUsed()), "minimax.sqlite"))
	if err := m.db.DB().AutoMigrate(storage.NewHistory()); err != nil {
		return err
	}

	return err
}

func (m *MiniMax) Prompt(ctx context.Context) {
	command.Rl.SetPrompt("\033[31mgptchat->minimax»\033[0m ")
}

func (m *MiniMax) Chat(ctx context.Context, message string, historyData []*storage.HistoryData) (string, error) {
	var completionText string
	resp, err := m.client.CreateCompletionStream(ctx, &minimax.ChatCompletionRequest{
		Model:    minimax.Abab5Dot5,
		Messages: buildMessage(historyData, driver.Viper().GetString("models.minimax.username"), message),
	})
	if err != nil {
		return "", err
	}
	defer resp.Close()
	for {
		res, err := resp.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return "", err
		}
		if res.Choices[0].FinishReason == "" {
			fmt.Fprint(command.Rl.Stderr(), res.Choices[0].Messages[0].Text)
		}
		if res.Reply != "" {
			completionText = res.Reply
		}
	}

	return completionText, nil
}

func (m *MiniMax) SaveHistory(ctx context.Context, history storage.HistoryObj) error {
	return m.db.Insert(ctx, storage.NewHistory(storage.SetHistoryData(history.GetQuestion(), history.GetAnswer())))
}

func (m *MiniMax) GetHistory(ctx context.Context) ([]*storage.HistoryData, error) {
	return m.db.Get(ctx, 5)

}

var (
	_ driver.Driver      = (*MiniMax)(nil)
	_ driver.GPTTerminal = (*MiniMax)(nil)
)
