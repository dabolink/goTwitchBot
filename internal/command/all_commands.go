package command

func AllCommands() []Command {
	return []Command{
		NewCountCommand(),
		&TestCommand{},
	}
}
