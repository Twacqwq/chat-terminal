package storage

import (
	"context"
	"sort"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type storage struct {
	driver string
	source string
	conf   *gorm.Config
}

func newStorage(driver, source string, opts ...storageOptions) *storage {
	s := &storage{
		driver: driver,
		source: source,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

type storageOptions func(*storage)

func withStorageConfig(conf *gorm.Config) storageOptions {
	return func(s *storage) {
		s.conf = conf
	}
}

func (s *storage) openDB() (*gorm.DB, error) {
	switch s.driver {
	case "sqlite":
		return gorm.Open(sqlite.Open(s.source), s.conf)
	default:
		return nil, gorm.ErrUnsupportedDriver
	}

}

type SqliteStorage struct {
	*storage
	db *gorm.DB
}

func OpenSqliteStorage(source string, conf ...gorm.Config) SqliteRepo {
	c := &gorm.Config{}
	if len(conf) > 0 {
		c = &conf[0]
	}

	sqliteStorage := &SqliteStorage{
		storage: newStorage("sqlite", source, withStorageConfig(c)),
		db:      nil,
	}

	db, err := sqliteStorage.openDB()
	if err != nil {
		panic(err)
	}
	sqliteStorage.db = db

	return sqliteStorage
}

func (sl *SqliteStorage) Insert(ctx context.Context, data *HistoryData) error {
	return sl.db.Create(data).Error
}

func (sl *SqliteStorage) Get(ctx context.Context, limit int) ([]*HistoryData, error) {
	var historyData []*HistoryData
	err := sl.db.Order("id DESC").Limit(limit).Find(&historyData).Error
	sort.Slice(historyData, func(i, j int) bool {
		return historyData[i].ID < historyData[j].ID
	})
	return historyData, err
}

func (sl *SqliteStorage) ClearAll(ctx context.Context) error {
	// TODO clear data table
	return nil
}

func (sl *SqliteStorage) DB() *gorm.DB {
	return sl.db
}

var (
	_ Repo = (*SqliteStorage)(nil)
	_ Repo = (*storage)(nil)
)
