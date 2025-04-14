package parser

import (
	"encoding/json"
	"errors"
	"log/slog"
	"time"
)

var (
	ErrInvalidMessage = errors.New("invalid message")
)

const (
	MessageType_Invalid          = "invalid"
	MessageType_SessionWelcome   = "session_welcome"
	MessageType_SessionKeepalive = "session_keepalive"
	MessageType_Notification     = "notification"
)

type Metadata struct {
	MessageID        string    `json:"message_id"`
	MessageType      string    `json:"message_type"`
	MessageTimestamp time.Time `json:"message_timestamp"`
}

type MessagePayload interface {
	Raw() string
}

type Message struct {
	Metadata Metadata
	Payload  MessagePayload
}

type MessageWrapper struct {
	Metadata Metadata       `json:"metadata"`
	Payload  map[string]any `json:"payload"`
}

func Parse(input []byte) (Message, error) {

	metadata, payload, err := parseMessage(input)
	if err != nil {
		panic(err)
	}
	if metadata.MessageType != MessageType_SessionKeepalive {
		slog.Debug(string(input))
	}
	return Message{Metadata: metadata, Payload: payload}, nil

}

func parseMessage(raw []byte) (Metadata, MessagePayload, error) {
	var wrapper MessageWrapper
	err := json.Unmarshal(raw, &wrapper)
	if err != nil {
		return Metadata{}, nil, err
	}

	payload, err := parsePayload(wrapper.Metadata.MessageType, wrapper.Payload)
	return wrapper.Metadata, payload, err
}

func parsePayload(messageType string, in map[string]any) (MessagePayload, error) {
	switch messageType {
	case MessageType_SessionWelcome:
		return decodePayload(in, SessionWelcome{})
	case MessageType_SessionKeepalive:
		return decodePayload(in, SessionKeepAlive{})
	case MessageType_Notification:
		return decodePayload(in, Notification{})
	default:
		slog.Error("failed to parse message", slog.Attr{Key: "messageType", Value: slog.StringValue(messageType)}, slog.Attr{Key: "payload", Value: slog.AnyValue(in)})
		return nil, errors.New("couldn't parse message of type: " + messageType)
	}
}

func decodePayload[T any](input map[string]any, output T) (t T, err error) {
	b, err := json.Marshal(input)
	if err != nil {
		return
	}
	json.Unmarshal(b, &output)
	return output, err
}
