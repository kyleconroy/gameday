package reddit

import (
  "net/http"
  "encoding/json"
  "io"
)

type Client struct {
  Username, Password, Hash  string
  hclient                    http.Client
}

func New(u string, p string) *Client {
  c := Char{Username: u, Password: p}
  payload = {'user': u, 'passwd': p}
  c.login(payload)
  return &c
}

func (c *Client) login () {
  c.hclient = &http.Client{}

  c.Post('https://ssl.reddit.com/api/login')  
}

func (c *Client) Post (url string) {
  return c.generate_request("POST", url, nil)
}


func (c *Client) Get (url string) {
  return c.generate_request("GET", url, nil)
}

func (c *Client) generate_request (t string, url string, body io.Reader) {
  req, err := http.NewRequest(t, url, body)
  if (err == nil) {
    return req
  } else {
    return err
  }
}