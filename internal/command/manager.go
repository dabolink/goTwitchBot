package command

import "errors"

var (
	ErrCmdNotFound = errors.New("cmd not found")
	DefaultManager = NewManager(NewCommandStore(AllCommands()))
)

type CommandStore interface {
	Get() []Command
	Save(command Command)
	Delete(command Command)
}

type CommandManager struct {
	commandStore   CommandStore
	loadedCommands []Command
}

func (manager *CommandManager) Process(info CommandInfo) {
	for _, cmd := range manager.loadedCommands {
		cmd.Process(info)
	}
}

func (manager *CommandManager) Init(store CommandStore) {
	manager.loadedCommands = store.Get()

}

func (manager *CommandManager) Delete(command Command) {
	remaining := make([]Command, 0, len(manager.loadedCommands)-1)
	for _, cmd := range manager.Get() {
		if cmd == command {
			continue
		}
		remaining = append(remaining, cmd)
	}
	manager.loadedCommands = remaining
}

func (manager *CommandManager) Add(command Command) {
	manager.loadedCommands = append(manager.loadedCommands, command)
}

func (manager *CommandManager) Get() []Command {
	return manager.loadedCommands
}

func NewManager(commandStore CommandStore) *CommandManager {
	manager := &CommandManager{
		commandStore: commandStore,
	}
	manager.Init(commandStore)
	return manager
}
