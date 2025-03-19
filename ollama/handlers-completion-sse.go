package ollama

import (
  "bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
func (c *Client) GenerateCompletionSSE(ctx context.Context, req GenerateCompletionReq) (<-chan CompletionSSERes, <-chan  error) {
  if !req.Stream {
    errCh := make(chan error, 1)
    errCh <- fmt.Errorf("use GenerteCompletion for stream disabled req")
    return nil, errCh
  }

  errCh := make(chan error, 1)
  outCh := make(chan CompletionSSERes)

  go func ()  {
    defer close(errCh)
    defer close(outCh)
    jsonData, err := json.Marshal(req)
    if err != nil {
      errCh <- fmt.Errorf("failed to marshal req: %w", err)
      return 
    }
    resp, err := c.makeHttpRequest(ctx, http.MethodPost, c.BaseUrl + GENERATE_COMPLETION_ENDPOINT, jsonData)
    if err != nil {
      errCh <- fmt.Errorf("failed to make http request: %w", err)
      return
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
      body, _ := io.ReadAll(resp.Body)
      errCh <- fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
      return
    }

    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
      line := scanner.Text()
      var chunk CompletionSSERes
      if err := json.Unmarshal([]byte(line), &chunk); err != nil {
        continue
      }
      outCh <- chunk
      if chunk.Done {
        break
      }
    }

    if err := scanner.Err(); err != nil {
      errCh <- fmt.Errorf("error reading stream: %w", err)
    }

  }()

  return outCh, errCh 
}

// Server-Side-Event driven
func (c *Client) GenerateChatCompletionSSE(ctx context.Context, req GenerateChatCompletionReq) (<-chan ChatCmpletionSSERes, <-chan error) {
  if !req.Stream {
    errCh := make(chan error,1)
    errCh <- fmt.Errorf("use GenerteChatCompletion for stream disabled req")
    close(errCh)
    return nil, errCh
  }

  errCh := make(chan error,1)
  outCh := make(chan ChatCmpletionSSERes)
  go func ()  {
    defer close(errCh)
    defer close(outCh)

    jsonData, err := json.Marshal(req)
    if err != nil {
      errCh <- fmt.Errorf("failed to marshal req: %w",err)
      return     
    }
    resp, err := c.makeHttpRequest(ctx, http.MethodPost, c.BaseUrl + GENERATE_CHAT_COMPLETION_ENDPOINT, jsonData)
    if err != nil {
      errCh <- fmt.Errorf("failed to make http request req: %w",err)
      return     
    }
    defer resp.Body.Close()
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan(){
      line := scanner.Text()
      var chunk ChatCmpletionSSERes 
      if err := json.Unmarshal([]byte(line), &chunk); err != nil {
        continue
      }
      outCh <- chunk
      if  chunk.Done {
        break
      }
    }

    if err := scanner.Err(); err != nil {
      errCh <- fmt.Errorf("error reading stream: %w", err)
      return
    }

  }()

  return outCh, errCh
 
}

