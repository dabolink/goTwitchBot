package command

import "github.com/google/uuid"

func AllCommands() []Command {
	return []Command{
		NewCountCommand(),
		&MultiCommand{
			id:       uuid.New(),
			Commands: make(map[string]Command),
		},
		&EchoCommand{},
	}
}
