package ollama

import "net/http"

type Client struct{
  BaseUrl string
  httpC *http.Client
}

func NewClient(baseUrl string) *Client {
  var httpClient = &http.Client{}
  return &Client{
    BaseUrl: baseUrl,
    httpC : httpClient,
  }
}
