package driver

import (
	"context"

	"github.com/Twacqwq/gpt-terminal/internal/storage"
)

type Driver interface {
	Meta
}

type Meta interface {
	// Init client struct
	Init(context.Context) error
	// Prompt set prompt style
	Prompt(context.Context)
}

type GPTTerminal interface {
	Driver

	// Chat is a conversation with the model
	Chat(ctx context.Context, messages string, historyData []*storage.HistoryData) (string, error)
	// SaveHistory Save the content of each conversation
	SaveHistory(ctx context.Context, history storage.HistoryObj) error
	// GetHistory Get the latest records
	GetHistory(ctx context.Context) ([]*storage.HistoryData, error)
}
