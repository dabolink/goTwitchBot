package main

import (
	"fmt"
	"goWebsocket/internal/env"
	"goWebsocket/internal/twitch/eventsub/helix"
)

type Config struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectUri  string `json:"redirect_uri"`
}

func main() {
	var config Config
	err := env.Load(&config)
	if err != nil {
		panic(err)
	}
	resp, err := helix.GetToken(config.ClientID, config.ClientSecret, config.Code, config.RedirectUri)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.AccessToken)
}
