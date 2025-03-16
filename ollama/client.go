package ollama

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



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



func (c *Client) Generate(ctx context.Context, req GenerateRequeset) (*GenerateResponse, error) {
  if req.Stream == true {
    return nil, fmt.Errorf("use GenerteStream for stream enabled req")
  }
  jsonData, err := json.Marshal(req)
  if err != nil {
    return nil, fmt.Errorf("failed to marshal req: %w", err)
  }
  resp, err := c.makeHttpRequest(ctx, http.MethodPost, c.BaseUrl + GENERATE_ENDPOINT, jsonData)
  if err != nil {
    return nil, fmt.Errorf("failed to make http request: %w", err) 
  }
  defer resp.Body.Close()

  if resp.StatusCode >= 400 {
    body, _ := io.ReadAll(resp.Body)
    return nil, fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
  }

  // Decode JSON response
  var gr GenerateResponse
  if err := json.NewDecoder(resp.Body).Decode(&gr); err != nil {
    return nil, fmt.Errorf("failed to decode response: %w", err)
  }

  return &gr, nil

}

func (c *Client) GenerateStreamSSE(ctx context.Context, req GenerateRequeset, handler func(StreamResponse)) error {
  if req.Stream == false {
    return fmt.Errorf("use Generte for stream disabled req")
  }
  jsonData, err := json.Marshal(req)
  if err != nil {
    return fmt.Errorf("failed to marshal req: %w", err)
  }
  resp, err := c.makeHttpRequest(ctx, http.MethodPost, c.BaseUrl + GENERATE_ENDPOINT, jsonData)
  if err != nil {
    return fmt.Errorf("failed to make http request: %w", err) 
  }
  defer resp.Body.Close()

  if resp.StatusCode >= 400 {
    body, _ := io.ReadAll(resp.Body)
    return fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
  }

  fmt.Println("Reading JSON stream...")

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

	fmt.Println("Stream processing finished.")
	return nil

}



