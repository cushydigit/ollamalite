package ollama

const (
  API = "/api"
  GENERATE_COMPLETION_ENDPOINT = API + "/generate"
  GENERATE_CHAT_COMPLETION_ENDPOINT = API + "/chat"
  CREATE_MODEL_MODEL = API + "/create"
  LIST_LOCALS_MODLES_ENDPOINT = API + "/tags"
  LIST_RUNNING_MODLES_ENDPOINT = API + "/ps"
  // TODO: check /api/embed
  GENERATE_EMBEDDING_ENDPOINT = API + "/embeddings"
)

type GenerateCompletionReq struct {  
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

type GenerateCompletionRes struct {  
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

type GenerateChatCompletionReq struct {
  Model string `json:"model"`
  Messages []Message `json:"messages"`
  Tools []string `json:"tools"` 
  Format string `json:"format"`
  Stream bool  `json:"stream"`
  KeepAlive int `json:"keep_alive"`
}

type Message struct {
  Role string `json:"role"`
  Content string `json:"content"`
  Image []string `json:"image"`
  ToolCalls []string `json:"tool_calls"`
}

// TODO: add feilds
type GenerateChatCompletionRes struct {
  
}

type LoadModelRes struct {
  Model string `json:"model"`
  CreatedAt string `json:"created_at"`
  Response string `json:"response"`
  Done bool `json:"done"`
}

type UnloadModelRes struct {
  Model string `json:"model"`
  CreatedAt string `json:"created_at"`
  Response string `json:"response"`
  Done bool `json:"done"`
  DoneReason string `json:"done_reason"`
}

type ListRunningModelsRes struct {
  Models []RunningModel `json:"models"` 
}

type ListLocalModelsRes struct {
  Models []LocalModel `json:"models"`
}

type RunningModel struct {
  Name string `json:"name"`
  ModifiedAt string `json:"modified_at"`
  Size int64 `json:"size"`
  Digest string `json:"digest"` 
  Details Details `json:"details"`
  ExpireAt string `json:"expire_at"`
  SizeVram int64 `json:"size_vram"`
}

type LocalModel struct {
  Name string `json:"name"`
  ModifiedAt string `json:"modified_at"`
  Size int64 `json:"size"`
  Digest string `json:"digest"` 
  Details Details `json:"details"`
}   

type Details struct {
  Format string `json:"format"`
  Family string `json:"family"`
  Families []string `json:"famieies"` // handle null references
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
	TotalDuration   int64    `json:"total_duration,omitempty"`
	LoadDuration    int64    `json:"load_duration,omitempty"`
	EvalCount   int      `json:"eval_count,omitempty"`
	EvalTime    int64    `json:"eval_duration,omitempty"`
}

type StreamCallBack func(StreamResponse)

