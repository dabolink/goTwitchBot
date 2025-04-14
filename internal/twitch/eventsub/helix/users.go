package helix

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

type UserList struct {
	Data []User `json:"data"`
}

type User struct {
	ID              string
	Login           string
	DisplayName     string
	Type            string
	BroadcasterType string
	Description     string
	ProfileImageUrl string
	OfflineImageUrl string
	ViewCount       int
	CreatedAt       time.Time
}

func GetUser(appToken string, clientID string, login string) (User, error) {
	users, err := GetUsers(appToken, clientID, []string{login})
	if err != nil {
		return User{}, err
	}
	if len(users) == 0 {
		return User{}, errors.New("user not found")
	}
	return users[0], nil
}

func GetUsers(appToken string, clientID string, logins []string) ([]User, error) {
	uri, err := url.Parse(baseUrl + "/users")
	if err != nil {
		return nil, err
	}
	query := uri.Query()
	for _, login := range logins {
		query.Add("login", login)
	}
	uri.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(context.Background(), "GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := makeRequest(req, appToken, clientID)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("failed to read response body")
	}
	var users UserList
	err = json.Unmarshal(b, &users)
	if err != nil {
		return nil, err
	}
	if len(users.Data) == 0 {
		logResponse("get user", resp)
		return nil, errors.New("no users")
	}
	return users.Data, nil
}
