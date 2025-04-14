package helix

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const authBaseUrl = "https://id.twitch.tv/oauth2"

type TokenResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int      `json:"expires_in"`
	Scope        []string `json:"scope"`
	TokenType    string   `json:"token_type"`
}

func GetToken(clientID string, clientSecret string, code string, redirectUri string) (TokenResponse, error) {
	uri, err := url.Parse(authBaseUrl + "/token")
	query := uri.Query()
	query.Add("client_id", clientID)
	query.Add("client_secret", clientSecret)
	query.Add("code", code)
	query.Add("grant_type", "authorization_code")
	query.Add("redirect_uri", redirectUri)
	uri.RawQuery = query.Encode()
	if err != nil {
		return TokenResponse{}, err
	}
	req, err := http.NewRequestWithContext(context.Background(), "POST", uri.String(), nil)
	if err != nil {
		return TokenResponse{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return TokenResponse{}, err
	}
	if !isOk(resp.StatusCode) {
		logResponse("invalid token response", resp)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return TokenResponse{}, err
	}
	var tokenResponse TokenResponse
	err = json.Unmarshal(b, &tokenResponse)
	return tokenResponse, err
}
