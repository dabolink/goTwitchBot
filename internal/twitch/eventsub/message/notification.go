package parser

import "time"

type Subscription struct {
	ID        string         `json:"id"`
	Status    string         `json:"status"`
	Type      string         `json:"type"`
	Condition map[string]any `json:"condititon"`
	Transport map[string]any `json:"transport"`
	CreatedAt time.Time      `json:"created_at"`
	Cost      int            `json:"cost"`
}
type Event struct {
	BroadcasterUserId           string              `json:"broadcaster_user_id"`
	BroadcasterUserLogin        string              `json:"broadcaster_user_login"`
	BroadcasterUserName         string              `json:"broadcaster_user_name"`
	ChatterUserId               string              `json:"chatter_user_id"`
	ChatterUserLogin            string              `json:"chatter_user_login"`
	ChatterUserName             string              `json:"chatter_user_name"`
	MessageId                   string              `json:"message_id"`
	Message                     NotificationMessage `json:"message"`
	MessageType                 string              `json:"message_type"`
	Badges                      []Badge             `json:"badges"`
	Cheer                       Cheer               `json:"cheer"`
	Color                       string              `json:"color"`
	Reply                       Reply               `json:"reply"`
	ChannelPointsCustomRewardID string              `json:"channel_points_custom_reward_id"`
	SourceBroadcasterUserId     string              `json:"source_broadcaster_user_id"`
	SourceBroadcasterUserName   string              `json:"source_broadcaster_user_name"`
	SourceBroadcasterUserLogin  string              `json:"source_broadcaster_user_login"`
	SourceMessageID             string              `json:"source_message_id"`
	SourceBadges                []Badge             `json:"source_badges"`
}

type Reply struct {
	ParentMessageId   string `json:"parent_message_id"`
	ParentMessageBody string `json:"parent_message_body"`
	ParentUserId      string `json:"parent_user_id"`
	ParentUserName    string `json:"parent_user_name"`
	ParentUserLogin   string `json:"parent_user_login"`
	ThreadMessageID   string `json:"thread_message_id"`
	ThreadUserID      string `json:"thread_user_id"`
	ThreadUserName    string `json:"thread_user_name"`
	ThreadUserLogin   string `json:"thread_user_login"`
}

type Cheer struct {
	Bits int `json:"bits"`
}

type Badge struct {
	SetID string `json:"set_id"`
	ID    string `json:"id"`
	Info  string `json:"info"`
}

type Fragment struct {
	Type      string    `json:"type"`
	Text      string    `json:"text"`
	Cheermote Cheermote `json:"cheermote"`
	Emote     Emote     `json:"emote"`
	Mention   Mention   `json:"mention"`
}

type Mention struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserLogin string `json:"user_login"`
}

type Emote struct {
	ID         string   `json:"id"`
	EmoteSetID string   `json:"emote_set_id"`
	OwnerID    string   `json:"owner_id"`
	Format     []string `json:"format"`
}

type Cheermote struct {
	Prefix string `json:"prefix"`
	Bits   int    `json:"bits"`
	Tier   int    `json:"tier"`
}

type NotificationMessage struct {
	Text      string     `json:"text"`
	Fragments []Fragment `json:"fragments"`
}

type Notification struct {
	Subscription Subscription `json:"subscription"`
	Event        Event        `json:"event"`
}

func (Notification) Raw() string {
	return ""
}
