package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/shahinrahimi/ollamalite/ollama"
)

var prompts = []string{
  "Describe humans in one words",
  "What's the secret to infinite wealth?",
}


func generateCompletionExample() {
 // creating ollamalite client
  oc := ollama.NewClient("http://localhost:11434")
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

  fmt.Printf("Response: %s\n", resp.Response)

}


func generateCompletionStreamExample() {
   // creating ollamalite client
  oc := ollama.NewClient("http://localhost:11434")
  // creating request
  req := ollama.GenerateCompletionReq{
    Model: "llama3.2:latest",
    Prompt: "Describe humans in one words.",
    Stream: true,
  }

  // make a stream json response
  if err := oc.GenerateCompletionStream(context.TODO(), req, func(resp ollama.StreamResponse){
    // fmt.Println("Received response chunk: ", resp)
    fmt.Print(resp.Response)
    // os.Stdout.Sync()
  }); err != nil {
    log.Fatal(err)
  }

}


func main() {
  generateCompletionExample()
  generateCompletionStreamExample()
  log.Fatal(errors.New("that's just happened"))

}
