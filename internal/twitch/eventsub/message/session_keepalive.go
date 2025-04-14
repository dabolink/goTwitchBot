package parser

type SessionKeepAlive struct{}

// Raw implements MessagePayload.
func (w SessionKeepAlive) Raw() string {
	return ""
}
