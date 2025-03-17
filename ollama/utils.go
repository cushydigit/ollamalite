package ollama

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func (c *Client) makeHttpRequest(ctx context.Context, method, url string, requestBody any) (*http.Response, error) {
  
  var body io.Reader
  if requestBody != nil {
    switch v := requestBody.(type) {
    case []byte:
      body = bytes.NewReader(v)
    case string:
      body = bytes.NewReader([]byte(v))
    default:
      jsonData, err := json.Marshal(v)
      if err != nil {
        return nil, fmt.Errorf("failed to marshal request body: %w", err)
      }
      body = bytes. NewReader(jsonData)
    }
  }

  req, err := http.NewRequestWithContext(ctx, method, url, body)
  if err != nil {
    return nil, err
  }

  req.Header.Set("Content-Type", "application/json")

  resp, err := c.httpC.Do(req) 
  if err != nil {
    return nil, err 
  }

  return resp, nil 
}

