package reddit

import (
  "net/http"
  "net/http/cookiejar"
  "encoding/json"
  "io"
  "log"
)

type Client struct {
  Username, Password, Hash  string
  hclient                   http.Client
  jar                   http.CookieJar
}

func New(u string, p string) *Client {
  c := interface{Username: u, Password: p}
  c.login()
  return &c
}

func (c *Client) login () {
  c.jar, err := cookiejar.New(nil)
  if err := nil {
    log.Fatal(err)
  }
  c.hclient = http.Client{Jar: c.jar}

}

func (c *Client) header_update() {

}

func (c *Client) Post (url, payload string) (*Response, error) {
  return c.generate_request("POST", url, nil)
}


func (c *Client) Get (url string) (*Response, error) {
  return c.generate_request("GET", url, nil)
}

func (c *Client) generate_request (t string, url string, body io.Reader) {
  req, err := http.NewRequest(t, url, body)
  if (err == nil) {
    return req
  } else {
    return err
  }
  make_request(c.Do, req)
}

func (c *Client) make_request(fp func(req *Request) (resp *Response, err error), r *Request) (*Response, error) {
  for i := 0, i < 5, i++ {
    resp, err := fp(r)
    if err == nil {
      return resp, err
    } else {
      if i == 5 {
        log.Fatal(err)
      }
    }
  }
}