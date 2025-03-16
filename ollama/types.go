package ollama

const (
  API = "/api"
  GENERATE_ENDPOINT = API + "/generate"
  GENERATE_CHAT_ENDPOINT = API + "/chat"
  LIST_MODLES_ENDPOINT = API + ""
)

type GenerateRequeset struct {  
  Model string `json:"model"`
  Prompt string `json:"prompt"`   
  Stream bool `json:"stream"` 
  Suffix string `json:"suffix,omitempty"`
  Image []string `json:"image,omitempty"` 
  Format string `json:"format,omitempty"` 
  System string `json:"system,omitempty"` // 
  Raw bool `json:"raw,omitempty"` // default false
  Options Options `json:"options,omitempty"`
  KeepAlive string `json:"keep_alive,omitempty"` 
}

type Options struct {
  Seed int64 `json:"seed,emitempty"`
  Temperature float64 `json:"temperature,omitempty"`
  NumCtx int `json:"num_ctx,omitempty"` 
}

type GenerateResponse struct {  
  Model string `json:"model"`
  Response string `json:"response"`
  Done bool `json:"done"` 
  PromptEvalCount int64  `json:"prompt_eval_count"`
  PromptEvalDuration int64 `json:"prompt_eval_duration"`
  EvalCount int64 `json:"eval_count"`
  EvalDuration int64 `json:"eval_duration"`
  TotalDuration int64 `json:"total_duration"`
  LoadDuration int64 `json:"load_duration"`
}

type GenerateChatRequest struct {
  Model string `json:"model"`
  Messages []Message `json:"messages"`
  Tools []string `json:"tools"` 
  Format string `json:"format"`
  Stream bool  `json:"stream"`
  KeepAlive string `json:"keep_alive"`
}

type Message struct {
  Role string `json:"role"`
  Content string `json:"content"`
  Image []string `json:"image"`
  ToolCalls []string `json:"tool_calls"`
}

// TODO: add feilds
type GenerateChatResponse struct {
  
}

type ListModelsResponse struct {
  Models []Model `json:"models"`
}

type Model struct {
  Name string `json:"name"`
  ModifiedAt string `json:"modified_at"`
  Size int64 `json:"size"`
  Digest string `json:"digest"` 
  Details Details `json:"details"`
}   

type Details struct {
  Format string `json:"format"`
  Family string `json:"family"`
  Families *string `json:"famieies"` // handle null references
  ParamaterSize string `json:"paramater_size"`
  QuantizationLevel string `json:"quantization_level"`
}


type StreamResponse struct {
	Model       string   `json:"model"`
	CreatedAt   string   `json:"created_at"`
	Response    string   `json:"response"` // The actual text content
	Done        bool     `json:"done"`
	DoneReason  string   `json:"done_reason,omitempty"`
	Context     []int    `json:"context,omitempty"`
	TotalTime   int64    `json:"total_duration,omitempty"`
	LoadTime    int64    `json:"load_duration,omitempty"`
	EvalCount   int      `json:"eval_count,omitempty"`
	EvalTime    int64    `json:"eval_duration,omitempty"`
}

type StreamCallBack func(StreamResponse)

