package minimax

import (
	"github.com/Twacqwq/go-minimax"
	"github.com/Twacqwq/gpt-terminal/internal/storage"
)

func miniMaxClient(apiToken, groupId string) (*minimax.Client, error) {
	return minimax.NewClient(apiToken, groupId), nil
}

func buildMessage(historyData []*storage.HistoryData, senderName string, contents ...string) (messages []minimax.Message) {
	for _, v := range historyData {
		messages = append(messages, minimax.Message{
			SenderType: minimax.ChatMessageRoleUser,
			SenderName: senderName,
			Text:       v.Q,
		})
		messages = append(messages, minimax.Message{
			SenderType: minimax.ChatMessageRoleBot,
			SenderName: minimax.ModelBot,
			Text:       v.A,
		})
	}

	for _, content := range contents {
		messages = append(messages, minimax.Message{
			SenderType: minimax.ChatMessageRoleUser,
			SenderName: senderName,
			Text:       content,
		})
	}

	return
}
