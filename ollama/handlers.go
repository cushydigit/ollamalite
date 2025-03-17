package ollama

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocalModels(ctx context.Context) (*ListLocalModelsRes, error) {
  resp, err := c.makeHttpRequest(ctx, http.MethodGet, c.BaseUrl + LIST_LOCALS_MODLES_ENDPOINT, nil)
  if err != nil {
    return nil, fmt.Errorf("failed to make http request: %w", err)
  }
  defer resp.Body.Close()

  if resp.StatusCode >= 400 {
    body, _ := io.ReadAll(resp.Body)
    return nil, fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
  }

  var res ListLocalModelsRes
  if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
    return nil, fmt.Errorf("failed to decode response: %w", err)
  }

  return &res, nil

}
func (c *Client) ListRunningModels(ctx context.Context) (*ListRunningModelsRes, error) {
  resp, err := c.makeHttpRequest(ctx, http.MethodGet, c.BaseUrl + LIST_RUNNING_MODLES_ENDPOINT, nil)
  if err != nil {
    return nil, fmt.Errorf("failed to make http request: %w", err)
  }
  defer resp.Body.Close()

  if resp.StatusCode >= 400 {
    body, _ := io.ReadAll(resp.Body)
    return nil, fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
  }

  var res ListRunningModelsRes
  if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
    return nil, fmt.Errorf("failed to decode response: %w", err)
  }

  return &res, nil

}


func (c *Client) UnLoadModel(ctx context.Context, modelName string) (*UnloadModelRes, error){
  var req GenerateChatCompletionReq
  req.Model = modelName
  req.KeepAlive = 0
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
  var res UnloadModelRes 
  if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
    return nil, fmt.Errorf("failed to decode response: %w", err)
  }

  return &res, nil

}


func (c *Client) LoadModel(ctx context.Context, modelName string) (*LoadModelRes, error){
  var req GenerateChatCompletionReq
  req.Model = modelName
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
  var res LoadModelRes 
  if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
    return nil, fmt.Errorf("failed to decode response: %w", err)
  }

  return &res, nil

}
func (c *Client) GenerateCompletion(ctx context.Context, req GenerateCompletionReq) (*GenerateCompletionRes, error) {
  if req.Stream == true {
    return nil, fmt.Errorf("use GenerteStreamSSE for stream enabled req")
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
  if req.Stream == true {
    return nil, fmt.Errorf("use GenerteChatStreamSSE for stream enabled req")
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
// SSE
func (c *Client) GenerateCompletionStream(ctx context.Context, req GenerateCompletionReq, handler func(StreamResponse)) error {
  if req.Stream == false {
    return fmt.Errorf("use Generte for stream disabled req")
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

		var chunk StreamResponse
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

func (c *Client) GenerateCompletionChatStreamSSE(ctx context.Context, req GenerateChatCompletionReq, handler func(StreamResponse)) error {
  if req.Stream == false {
    return fmt.Errorf("use GenerteChat for stream disabled req")
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

		var chunk StreamResponse
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



