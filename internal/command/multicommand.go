package command

import (
	"strings"

	"github.com/google/uuid"
)

type MultiCommand struct {
	id       uuid.UUID
	Commands map[string]Command
}

func (cmd *MultiCommand) ID() uuid.UUID {
	return cmd.id
}

func (cmd *MultiCommand) matches(cmdInfo CommandInfo) bool {
	return strings.HasPrefix(cmdInfo.MessageInfo.Text, cmdInfo.BotInfo.CmdPrefix)

}

func (cmd *MultiCommand) Process(cmdInfo CommandInfo) {

	commandParts := strings.Split(cmdInfo.MessageInfo.Text, " ")
	commandName := commandParts[0][1:]

	subCmd, ok := cmd.Commands[commandName]
	if !ok {
		return
	}

	subCmd.Process(cmdInfo)
}
