
package ollama

import (
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

