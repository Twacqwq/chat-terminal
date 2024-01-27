package storage

import (
	"context"

	"gorm.io/gorm"
)

type HistoryData struct {
	ID uint64 `gorm:"primaryKey"`
	Q  string
	A  string
}

func NewHistory(opts ...dataOpts) *HistoryData {
	historyData := &HistoryData{}
	for _, opt := range opts {
		opt(historyData)
	}
	return historyData
}

type dataOpts func(*HistoryData)

func SetHistoryData(q, a string) dataOpts {
	return func(hd *HistoryData) {
		hd.Q = q
		hd.A = a
	}
}

type Repo interface {
	openDB() (*gorm.DB, error)
}

type SqliteRepo interface {
	Repo

	DB() *gorm.DB
	Insert(ctx context.Context, data *HistoryData) error
	Get(ctx context.Context, limit int) ([]*HistoryData, error)
	ClearAll(ctx context.Context) error
}

type HistoryObj interface {
	GetQuestion() string
	GetAnswer() string
}

type HistoryReader struct {
	question string
	answer   string
}

func NewHistoryReader(q, a string) *HistoryReader {
	return &HistoryReader{
		question: q,
		answer:   a,
	}
}

func (h *HistoryReader) GetQuestion() string {
	return h.question
}

func (h *HistoryReader) GetAnswer() string {
	return h.answer
}
