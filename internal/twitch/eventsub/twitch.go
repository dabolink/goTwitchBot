package twitch

import (
	"context"
	"errors"
	"goWebsocket/internal/maps"
	"goWebsocket/internal/twitch/eventsub/helix"
	parser "goWebsocket/internal/twitch/eventsub/message"
	ws "goWebsocket/internal/ws"
	"log/slog"
)

type ChannelBotBuilder interface {
	Build(key string) MessageReceiver
}

type Config struct {
	EnableSSL bool
	Pass      string
	Nick      string
	ClientID  string
	Channels  []string
}

type MessageReceiver interface {
	OnMessageReceive(metadata parser.Metadata, payload parser.Notification)
}

type ChannelMap interface {
	Get(id string) MessageReceiver
}

type Client struct {
	config         Config
	wsCli          *ws.Client
	channelManager ChannelMap
	messageCh      chan parser.Message
	ctx            context.Context
}

func (cli *Client) Run() any {
	slog.Debug("Starting twitch irc client...")
	go messageReader(cli)
	go messageReceiver(cli)
	return <-cli.ctx.Done()
}

func messageReader(cli *Client) {
	for {
		_, msg, err := cli.ReadMessage()
		if err != nil {
			panic(err)
		}
		message, err := parser.Parse(msg)
		if err != nil {
			panic(err)
		}
		cli.messageCh <- message
	}
}

func (cli *Client) ReadMessage() (int, []byte, error) {
	return cli.wsCli.ReadMessage()
}

func messageReceiver(cli *Client) {
	for v := range cli.messageCh {
		switch payload := v.Payload.(type) {
		case parser.SessionWelcome:
			botUser, err := helix.GetUser(cli.config.Pass, cli.config.ClientID, cli.config.Nick)
			if err != nil {
				panic(err)
			}
			broadcastUsers, err := helix.GetUsers(cli.config.Pass, cli.config.ClientID, cli.config.Channels)
			if err != nil {
				panic(err)
			}
			for _, broadcastUser := range broadcastUsers {
				resp, err := helix.CreateEventSubSubscription(cli.config.Pass, cli.config.ClientID, helix.CreateEventSubSubscriptionRequest{
					Type:    "channel.chat.message",
					Version: "1",
					Condition: map[string]any{
						"broadcaster_user_id": broadcastUser.ID,
						"user_id":             botUser.ID,
					},
					Transport: helix.Transport{
						Method:    "websocket",
						SessionID: payload.Session.ID,
					},
				})
				if err != nil {
					panic(err)
				}
				slog.Debug("session setup", slog.Attr{Key: "response", Value: slog.AnyValue(resp)})
			}
		case parser.Notification:
			receiver := cli.channelManager.Get(payload.Event.BroadcasterUserId)
			if receiver != nil {
				receiver.OnMessageReceive(v.Metadata, payload)
			}
		}

	}
}

func NewClient(config Config, botBuilder ChannelBotBuilder) (*Client, error) {
	url := "wss://eventsub.wss.twitch.tv/ws"
	if config.EnableSSL {
		return nil, errors.New("SSL not supported")
	}

	ctx := context.Background()
	cli, err := ws.NewClient(url, ctx)

	return &Client{
		config:         config,
		wsCli:          cli,
		ctx:            ctx,
		channelManager: maps.NewDefaultMap(botBuilder),
		messageCh:      make(chan parser.Message),
	}, err
}
