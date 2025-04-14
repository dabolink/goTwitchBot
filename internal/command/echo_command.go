package command

import (
	"fmt"
	"goWebsocket/internal/variables"

	"github.com/google/uuid"
)

type EchoCommand struct {
	id uuid.UUID
}

func (cmd *EchoCommand) ID() uuid.UUID {
	return cmd.id
}

func (cmd *EchoCommand) matches(cmdInfo CommandInfo) bool {
	return true
}

func (cmd *EchoCommand) Process(cmdInfo CommandInfo) {
	if !cmd.matches(cmdInfo) {
		return
	}
	updatedVar := cmdInfo.Variables.Update("echo", variables.Increment(1))
	cmdInfo.Logger.Info("echo variable update", "variable", "echo", "value", fmt.Sprintf("%v", updatedVar))
}
