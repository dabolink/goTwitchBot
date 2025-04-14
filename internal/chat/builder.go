package chat

import (
	"goWebsocket/internal/command"
	twitch "goWebsocket/internal/twitch/eventsub"
)

type SimpleChatBotBuilder struct {
	Config    Config
	Processor Processor
	Variables command.VariableManager
}

func (builder *SimpleChatBotBuilder) Build(key string) twitch.MessageReceiver {
	return NewBot(builder.Config, builder.Processor, builder.Variables)
}

func NewBotBuilder(config Config, processor Processor, variables command.VariableManager) *SimpleChatBotBuilder {
	return &SimpleChatBotBuilder{
		Config:    config,
		Processor: processor,
		Variables: variables,
	}
}
