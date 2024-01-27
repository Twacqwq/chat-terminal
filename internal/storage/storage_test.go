package storage_test

import (
	"testing"

	"github.com/Twacqwq/gpt-terminal/internal/storage"
)

type User struct {
	ID      uint64 `gorm:"primaryKey"`
	Name    string
	Address string
}

func TestStorage(t *testing.T) {
	s := storage.OpenSqliteStorage("./test.sqlite")
	user := &User{
		Name:    "222",
		Address: "3334",
	}
	s.DB().Create(&user)

}
