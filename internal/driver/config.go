package driver

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	v    *viper.Viper
)

func init() {
	once.Do(func() {
		v = viper.New()
		v.SetConfigFile(".gptchat")
		v.SetConfigType("json")
		v.AddConfigPath("$HOME")
	})
}

type Config struct {
	MiniMax *MiniMaxConf
}

type MiniMaxConf struct {
	AccessToken string
	GroupId     string
	UserName    string
}

func NewConfig() *Config {
	return &Config{
		MiniMax: &MiniMaxConf{},
	}
}

func Viper() *viper.Viper {
	return v
}
