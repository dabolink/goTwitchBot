package main

import (
	"goWebsocket/internal/chat"
	"goWebsocket/internal/command"
	"goWebsocket/internal/env"
	twitch "goWebsocket/internal/twitch/eventsub"
	"goWebsocket/internal/variables"
	"io"
	"log/slog"
	"os"
)

type Config struct {
	Password string   `json:"password"`
	Username string   `json:"username"`
	ClientID string   `json:"client_id"`
	Channels []string `json:"channels"`
}

func init() {
	logFile, err := os.OpenFile("log/twitch.jsonl", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	options := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(io.MultiWriter(logFile, os.Stdout), &options)))
}

func main() {
	var config Config
	err := env.Load(&config)
	if err != nil {
		panic(err)
	}

	commandStore := command.NewCommandStore(command.AllCommands())

	client, err := twitch.NewClient(twitch.Config{
		Pass:     config.Password,
		Nick:     config.Username,
		ClientID: config.ClientID,
		Channels: config.Channels,
	}, chat.NewBotBuilder(chat.Config{CmdPrefix: "!"}, command.NewManager(commandStore), variables.NewManager()))

	if err != nil {
		panic(err)
	}

	client.Run()
}
