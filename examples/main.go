package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/shahinrahimi/ollamalite/ollama"
)


type Application struct{
  oc *ollama.Client
}

var oc = ollama.NewClient("http://localhost:11434")

var prompts = []string{
  "Describe humans in one words",
  "What's the secret to infinite wealth?",
}


func generateCompletionExample() {
  // creating request
  req := ollama.GenerateCompletionReq{
    Model: "llama3.2:latest",
    Prompt: prompts[0],
    Stream: false,
  }
    
  // make a single json response
  resp, err := oc.GenerateCompletion(context.TODO(), req)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("\nResponse: %s\n", resp.Response)

}


func generateCompletionStreamExample() {
  // creating request
  req := ollama.GenerateCompletionReq{
    Model: "llama3.2:latest",
    Prompt: "Describe humans in one words.",
    Stream: true,
  }

  // make a stream json response
  if err := oc.GenerateCompletionSSE(context.TODO(), req, func(resp ollama.CompletionSSERes){
    fmt.Print(resp.Response)
  }); err != nil {
    log.Fatal(err)
  }

}

func generateChatCompletionExample() {
  req := ollama.GenerateChatCompletionReq{
    Model: "llama3.2:latest",
    Messages: []ollama.Message{
      {Role: "system", Content: "You are helpfull assisstant. providing wity and dry responses."},
      {Role: "user", Content: "Tell me an intersting space fact."},
    },
    Stream: false,
  }

  resp, err := oc.GenerateChatCompletion(context.TODO(), req)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("\n%s: %s\n", resp.Message.Role ,resp.Message.Content)
}

func generateChatCompletionStreamExample() {
  req := ollama.GenerateChatCompletionReq{
    Model: "llama3.2:latest",
    Messages: []ollama.Message{
      {Role: "system", Content: "You are helpfull assisstant. providing wity and dry responses."},
      {Role: "user", Content: "Tell me an intersting space fact."},
    },
    Stream: true,
  }

  if err := oc.GenerateChatCompletionSSE(context.TODO(), req,func(sr ollama.ChatCmpletionSSERes) {
    fmt.Print(sr.Messge.Content)

  }); err != nil {
    log.Fatal(err)
  }
}


func main() {
  generateChatCompletionStreamExample()
  generateChatCompletionExample()
  log.Fatal(errors.New("that's just happened"))
  generateCompletionExample()
  generateCompletionStreamExample()

}
