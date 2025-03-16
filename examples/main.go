package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shahinrahimi/ollamalite/ollama"
)


func main() {
  // creating ollamalite client
  oc := ollama.NewClient("http://localhost:11434")
  // creating request
  req := ollama.GenerateRequeset{
    Model: "llama3.2:latest",
    Prompt: "explain interfaces in 10 most popular languages and dofferences!",
  }
    
  // make a single json response
  resp, err := oc.Generate(context.TODO(), req)
  if err != nil {
    log.Fatal(err)
  }


  fmt.Println("Response: ", resp.Response)


  // make a stream json response

}
