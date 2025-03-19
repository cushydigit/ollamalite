package ollama

import (
  "bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
// Server-Side-Event driven
func (c *Client) GenerateCompletionSSE(ctx context.Context, req GenerateCompletionReq, handler CompletionSSECallback) error {
  if req.Stream == false {
    return fmt.Errorf("use GenerteCompletion for stream disabled req")
  }
  jsonData, err := json.Marshal(req)
  if err != nil {
    return fmt.Errorf("failed to marshal req: %w", err)
  }
  resp, err := c.makeHttpRequest(ctx, http.MethodPost, c.BaseUrl + GENERATE_COMPLETION_ENDPOINT, jsonData)
  if err != nil {
    return fmt.Errorf("failed to make http request: %w", err) 
  }
  defer resp.Body.Close()

  if resp.StatusCode >= 400 {
    body, _ := io.ReadAll(resp.Body)
    return fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
  }

	// Use bufio scanner to read each JSON object line by line
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("Raw JSON: %s\n", line) // Debugging line

		var chunk CompletionSSERes
		if err := json.Unmarshal([]byte(line), &chunk); err != nil {
      fmt.Printf("Failed to parse JSON: %s\n", err)
			continue
		}

		// Send parsed response to callback
		handler(chunk)

		// Stop processing if stream is marked as done
		if chunk.Done {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading stream: %w", err)
	}

	return nil

}
// Server-Side-Event driven
func (c *Client) GenerateChatCompletionSSE(ctx context.Context, req GenerateChatCompletionReq, handler ChatCompletionSSECallback ) error {
  if req.Stream == false {
    return fmt.Errorf("use GenerteChatCompletion for stream disabled req")
  }
  jsonData, err := json.Marshal(req)
  if err != nil {
    return fmt.Errorf("failed to marshal req: %w", err)
  }
  resp, err := c.makeHttpRequest(ctx, http.MethodPost, c.BaseUrl + GENERATE_CHAT_COMPLETION_ENDPOINT, jsonData)
  if err != nil {
    return fmt.Errorf("failed to make http request: %w", err) 
  }
  defer resp.Body.Close()

  if resp.StatusCode >= 400 {
    body, _ := io.ReadAll(resp.Body)
    return fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
  }

	// Use bufio scanner to read each JSON object line by line
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("Raw JSON: %s\n", line) // Debugging line

		var chunk ChatCmpletionSSERes 
		if err := json.Unmarshal([]byte(line), &chunk); err != nil {
			fmt.Printf("Failed to parse JSON: %s\n", err)
			continue
		}

		// Send parsed response to callback
		handler(chunk)

		// Stop processing if stream is marked as done
		if chunk.Done {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading stream: %w", err)
	}

	return nil

}

