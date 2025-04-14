package command

type CommandVisitor struct {
}

func (visitor *CommandVisitor) Visit(command Command, cmdInfo CommandInfo) {
	command.Process(cmdInfo)
}
