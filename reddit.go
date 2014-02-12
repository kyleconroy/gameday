package reddit

import (
  "net/http"
  "encoding/json"
)

type Client struct {
  Username  string
  Password  string
  Client    http.Client
  Hash      string
}

func (c *Client) login () {
  
}