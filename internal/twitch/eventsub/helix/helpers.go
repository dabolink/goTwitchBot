package helix

import (
	"io"
	"log/slog"
	"net/http"
)

func makeRequest(req *http.Request, appToken string, clientID string) (*http.Response, error) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+appToken)
	req.Header.Add("Client-Id", clientID)

	return http.DefaultClient.Do(req)
}

func logResponse(msg string, resp *http.Response) {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("error parsing body", slog.Attr{Key: "err", Value: slog.AnyValue(err)})
	}
	slog.Debug(msg, slog.Attr{Key: "uri", Value: slog.StringValue(resp.Request.RequestURI)}, slog.Attr{Key: "body", Value: slog.StringValue(string(b))})
}

func isOk(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 300
}
