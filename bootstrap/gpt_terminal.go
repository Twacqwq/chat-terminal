package bootstrap

import (
	"context"
	"fmt"

	"github.com/Twacqwq/gpt-terminal/drivers/minimax"
	"github.com/Twacqwq/gpt-terminal/internal/driver"
	"github.com/Twacqwq/gpt-terminal/internal/typs"
)

var ErrModelNotSupport = fmt.Errorf("model is not support")

func CreateTerminal(ctx context.Context, model string) (client driver.GPTTerminal, err error) {
	switch model {
	case typs.MiniMax:
		client = minimax.NewMiniMax()
		err = client.Init(ctx)
	default:
		err = ErrModelNotSupport
	}

	return
}
