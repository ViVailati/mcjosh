package models

import (
	client "github.com/ViVailati/mcjosh/http_client"
)

const (
	HttpClient = "http_client"
)

type Context struct {
	Data map[string]any
}

func (c *Context) GetHttpClient() *client.Client {
	return c.Data[HttpClient].(*client.Client)
}
