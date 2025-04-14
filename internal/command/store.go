package command

type InMemoryCommandStore struct {
	commands []Command
}

// Delete implements CommandStore.
func (i *InMemoryCommandStore) Delete(command Command) {
	panic("unimplemented")
}

func (i *InMemoryCommandStore) Get() []Command {
	return i.commands
}

// Save implements CommandStore.
func (i *InMemoryCommandStore) Save(command Command) {

}

func NewCommandStore(commands []Command) *InMemoryCommandStore {
	return &InMemoryCommandStore{
		commands: commands,
	}
}
