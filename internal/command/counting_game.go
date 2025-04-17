package command

import (
	"goWebsocket/internal/command/counter"
	"strconv"

	"github.com/google/uuid"
)

type CountingGameCommand struct {
	id      uuid.UUID
	counter Counter
}

type Counter interface {
	Play(string, int) error
}

func (cmd *CountingGameCommand) matches(cmdInfo CommandInfo) bool {
	_, ok := cmd.parse(cmdInfo)
	return ok
}

func (cmd *CountingGameCommand) parse(cmdInfo CommandInfo) (int, bool) {
	val, err := strconv.Atoi(cmdInfo.MessageInfo.Text)
	return val, err == nil
}

func (cmd *CountingGameCommand) ID() uuid.UUID {
	return cmd.id
}

func (cmd *CountingGameCommand) Process(cmdInfo CommandInfo) error {
	val, ok := cmd.parse(cmdInfo)
	if !ok {
		return nil
	}
	err := cmd.counter.Play(cmdInfo.MessageInfo.Chatter.ID, val)
	if err != nil {
		cmdInfo.Logger.Error("user failed counting game", "user", cmdInfo.MessageInfo.Chatter.DisplayName, "value", val, "err", err)
	}
	return nil
}

func NewCountCommand() *CountingGameCommand {
	return &CountingGameCommand{id: uuid.New(), counter: counter.NewCounter()}
}
