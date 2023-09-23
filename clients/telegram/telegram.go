package telegram

import "net/http"

type Client struct {
	host string
	basePath string
	client http.Client
}

func main(host string, token string) Client {
	return Client{
		host: host,
		basePath: newbasePath(token),
		client: http.Client{},
	}
}

func newbasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates() {
	
}

func (c *Client) SendMessage() {

}