package command

import (
	"fmt"
	"goWebsocket/internal/variables"

	"github.com/google/uuid"
)

type TestCommand struct {
	id uuid.UUID
}

func (cmd *TestCommand) ID() uuid.UUID {
	return cmd.id
}

func (cmd *TestCommand) matches(cmdInfo CommandInfo) bool {
	return true
}

func (cmd *TestCommand) Process(cmdInfo CommandInfo) error {
	if !cmd.matches(cmdInfo) {
		return nil
	}
	updatedVar := cmdInfo.Variables.Update("echo", variables.Increment(1))
	cmdInfo.Logger.Info("echo variable update", "variable", "echo", "value", fmt.Sprintf("%v", updatedVar))
	return nil
}
