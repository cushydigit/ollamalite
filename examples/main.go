package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/shahinrahimi/ollamalite/ollama"
)

func exampleGenerate() {
 // creating ollamalite client
  oc := ollama.NewClient("http://localhost:11434")
  // creating request
  req := ollama.GenerateRequeset{
    Model: "llama3.2:latest",
    Prompt: "explain interfaces in 10 most popular languages and dofferences!",
    Stream: false,
  }
    
  // make a single json response
  resp, err := oc.Generate(context.TODO(), req)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Response: %s", resp.Response)

}


func exampleGenerateStream() {
   // creating ollamalite client
  oc := ollama.NewClient("http://localhost:11434")
  // creating request
  req := ollama.GenerateRequeset{
    Model: "llama3.2:latest",
    Prompt: "please tell best movie of alltime. just one without explanation just name.",
    Stream: true,
  }

  // make a stream json response
  if err := oc.GenerateStreamSSE(context.TODO(), req, func(resp ollama.StreamResponse){
    //fmt.Println("Received response chunk: ", resp)
    fmt.Print(resp.Response)
    os.Stdout.Sync()
  }); err != nil {
    log.Fatal(err)
  }

  fmt.Println("\nStram ended.")

}


func main() {
  exampleGenerateStream()
  log.Fatal(errors.New("that's just happened"))
  exampleGenerate()

}
