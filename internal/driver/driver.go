package driver

import (
	"context"
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

	Chat(ctx context.Context, messages string) error
}
