package parser

import "time"

const (
	Status_Connected = "connected"
)

type SessionWelcome struct {
	Session Session
}

// Raw implements MessagePayload.
func (w SessionWelcome) Raw() string {
	return ""
}

type Session struct {
	ID                      string    `json:"id"`
	Status                  string    `json:"status"`
	KeepaliveTimeoutSeconds int       `json:"keepalive_timeout_seconds"`
	ReconnectUrl            string    `json:"reconnect_url"`
	ConnectedAt             time.Time `json:"recovery_url"`
}
