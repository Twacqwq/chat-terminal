package minimax

import (
	"github.com/Twacqwq/go-minimax"
)

func miniMaxClient(apiToken, groupId string) (*minimax.Client, error) {
	return minimax.NewClient(apiToken, groupId), nil
}
