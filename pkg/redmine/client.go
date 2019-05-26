package redmine

import "net/http"

type Client struct {
	endpoint string
	apikey   string
	*http.Client
	Limit  int
	Offset int
}

const (
	defaultLimit  int = -1
	defaultOffset int = -1
)

func NewClient(endpoint, apikey string) *Client {
	return &Client{endpoint, apikey, http.DefaultClient, defaultLimit, defaultOffset}
}

type Id struct {
	ID int `json:"id"`
}
type IdName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
