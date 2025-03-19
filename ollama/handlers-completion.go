package ollama

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
func (c *Client) GenerateCompletion(ctx context.Context, req GenerateCompletionReq) (*GenerateCompletionRes, error) {
  if req.Stream {
    return nil, fmt.Errorf("use GenerteCompletionSSE for stream enabled req")
  }
  
  jsonData, err := json.Marshal(req)
  if err != nil {
    return nil, fmt.Errorf("failed to marshal req: %w", err)
  }
  resp, err := c.makeHttpRequest(ctx, http.MethodPost, c.BaseUrl + GENERATE_COMPLETION_ENDPOINT, jsonData)
  if err != nil {
    return nil, fmt.Errorf("failed to make http request: %w", err) 
  }
  defer resp.Body.Close()

  if resp.StatusCode >= 400 {
    body, _ := io.ReadAll(resp.Body)
    return nil, fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
  }

  // Decode JSON response
  var res GenerateCompletionRes
  if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
    return nil, fmt.Errorf("failed to decode response: %w", err)
  }

  return &res, nil

}
   
func (c *Client) GenerateChatCompletion(ctx context.Context, req GenerateChatCompletionReq) (*GenerateChatCompletionRes, error) {
  if req.Stream {
    return nil, fmt.Errorf("use GenerteChatCompletionSSE for stream enabled req")
  }
  jsonData, err := json.Marshal(req)
  if err != nil {
    return nil, fmt.Errorf("failed to marshal req: %w", err)
  }
  resp, err := c.makeHttpRequest(ctx, http.MethodPost, c.BaseUrl + GENERATE_CHAT_COMPLETION_ENDPOINT, jsonData)
  if err != nil {
    return nil, fmt.Errorf("failed to make http request: %w", err) 
  }
  defer resp.Body.Close()

  if resp.StatusCode >= 400 {
    body, _ := io.ReadAll(resp.Body)
    return nil, fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
  }

  // Decode JSON response
  var res GenerateChatCompletionRes
  if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
    return nil, fmt.Errorf("failed to decode response: %w", err)
  }

  return &res, nil

}



