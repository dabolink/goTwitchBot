package chat

import (
	"goWebsocket/internal/command"
	parser "goWebsocket/internal/twitch/eventsub/message"
	"log/slog"
)

type Config struct {
	CmdPrefix string
}

type Counter interface {
	Play(chatterID string, val int) error
}

type Processor interface {
	Process(command.CommandInfo)
}

type Bot struct {
	config    Config
	processor Processor
	variables command.VariableManager
}

func (b *Bot) OnMessageReceive(metadata parser.Metadata, payload parser.Notification) {
	logger := slog.Default().With("broadcaster", payload.Event.BroadcasterUserName, "chatter", payload.Event.ChatterUserName, "text", payload.Event.Message.Text)

	info := b.buildCommandInfo(payload, logger)

	b.processor.Process(info)

	slog.Debug(metadata.MessageType, "message", payload)

}

func NewBot(config Config, processor Processor, variables command.VariableManager) *Bot {
	return &Bot{
		config:    config,
		processor: processor,
		variables: variables,
	}
}

func (b *Bot) buildCommandInfo(payload parser.Notification, logger *slog.Logger) command.CommandInfo {
	return command.CommandInfo{
		MessageInfo: command.MessageInfo{
			Logger: logger,
			Chatter: command.User{
				ID:          payload.Event.ChatterUserId,
				DisplayName: payload.Event.ChatterUserName,
				LoginName:   payload.Event.ChatterUserLogin,
			},
			Broadcaster: command.User{
				ID:          payload.Event.BroadcasterUserId,
				DisplayName: payload.Event.BroadcasterUserName,
				LoginName:   payload.Event.BroadcasterUserLogin,
			},
			Text: payload.Event.Message.Text,
		},
		Variables: b.variables,
		BotInfo: command.BotInfo{
			CmdPrefix: b.config.CmdPrefix,
		},
		Logger: logger,
	}
}
