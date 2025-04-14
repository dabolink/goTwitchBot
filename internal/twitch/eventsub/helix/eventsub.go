package helix

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const baseUrl = "https://api.twitch.tv/helix"

type CreateEventSubSubscriptionRequest struct {
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Condition any       `json:"condition"`
	Transport Transport `json:"transport"`
	ConduitID string    `json:"conduit_id,omitempty"`
}

type CreateEventSubSubscriptionResponse struct {
	Data []Subscription `json:"data"`
}

type Subscription struct {
	ID           string    `json:"id"`
	Status       string    `json:"status"`
	Type         string    `json:"type"`
	Version      string    `json:"version"`
	Condition    any       `json:"condition"`
	CreatedAt    time.Time `json:"created_at"`
	Transport    Transport `json:"transport"`
	ConnectedAt  time.Time `json:"connected_at"`
	ConduitID    string    `json:"conduit_id"`
	Cost         int       `json:"cost"`
	Total        int       `json:"total"`
	TotalCost    int       `json:"total_cost"`
	MaxTotalCost int       `json:"max_total_cost"`
}

type Transport struct {
	Method    string `json:"method"`
	Callback  string `json:"callback,omitempty"`
	SessionID string `json:"session_id,omitempty"`
	Secret    string `json:"secret,omitempty"`
}

func CreateEventSubSubscription(appToken string, clientID string, request CreateEventSubSubscriptionRequest) (CreateEventSubSubscriptionResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return CreateEventSubSubscriptionResponse{}, err
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", baseUrl+"/eventsub/subscriptions", bytes.NewReader(body))
	if err != nil {
		return CreateEventSubSubscriptionResponse{}, err
	}

	resp, err := makeRequest(req, appToken, clientID)
	if err != nil {
		return CreateEventSubSubscriptionResponse{}, nil
	}

	if !isOk(resp.StatusCode) {
		logResponse("create event sub subscription", resp)
		return CreateEventSubSubscriptionResponse{}, errors.New("invalid status code: " + fmt.Sprint(resp.StatusCode))
	}

	var responseBody CreateEventSubSubscriptionResponse
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return CreateEventSubSubscriptionResponse{}, nil
	}
	return responseBody, nil
}
