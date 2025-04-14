package ws

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	Ctx  context.Context
}

func NewClient(url string, ctx context.Context) (*Client, error) {
	requestHeader := make(http.Header)
	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, url, requestHeader)
	if err != nil {
		return nil, err
	}
	if !okResponseStatus(resp) {
		panic(*resp)
	}
	return &Client{conn: conn, Ctx: ctx}, nil
}

func (cli *Client) SendString(message string) error {
	return cli.conn.WriteMessage(1, []byte(message))
}

func (cli *Client) ReadMessage() (int, []byte, error) {
	messageType, bytes, err := cli.conn.ReadMessage()
	if err != nil {
		return -1, nil, err
	}
	return messageType, bytes, nil
}

func okResponseStatus(resp *http.Response) bool {
	return resp.StatusCode < 200 || resp.StatusCode > 299
}
