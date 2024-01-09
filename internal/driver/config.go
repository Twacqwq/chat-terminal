package driver

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	v    *viper.Viper
)

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
	once.Do(func() {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		v = viper.New()
		v.SetConfigFile(fmt.Sprintf("%s/.gptchat", homeDir))
		v.SetConfigType("json")
		v.AddConfigPath("$HOME")
	})

	return v
}
