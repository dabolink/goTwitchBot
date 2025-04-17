package command

import "github.com/google/uuid"

type Command interface {
	CommandMatcher
	CommandRunnable
	ID() uuid.UUID
}

type CommandMatcher interface {
	matches(CommandInfo) bool
}

type CommandRunnable interface {
	Process(CommandInfo) error
}
