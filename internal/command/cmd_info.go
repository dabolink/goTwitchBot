package command

import "goWebsocket/internal/variables"

type CommandInfo struct {
	MessageInfo MessageInfo
	BotInfo     BotInfo
	Variables   VariableManager
	Logger      Logger
}

type VariableManager interface {
	Set(name string, value variables.Variable)
	Get(name string) variables.Variable
	Update(name string, updateFn variables.VariableUpdateFn) variables.Variable
}

type BotInfo struct {
	CmdPrefix string
}

type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type User struct {
	ID          string
	DisplayName string
	LoginName   string
}

type MessageInfo struct {
	Logger      Logger
	Text        string
	Broadcaster User
	Chatter     User
}
