package main

import (
  "errors"
  "fmt"
  "net/http"
  "os"
  "github.com/graphql-go/handler"
)

func main() {
  graphqlHandler := handler.New(&handler.Config{
    Schema: &Schema, // comes from schema.go
    Pretty: true,
    Playground: true,
  })

  http.Handle("/graphql", graphqlHandler)

  fmt.Printf("starting server\n")

  err := http.ListenAndServe(":3333", nil)

  if errors.Is(err, http.ErrServerClosed) {
    fmt.Printf("server closed\n")
  } else if err != nil {
    fmt.Printf("error starting server: %s\n", err)
    os.Exit(1)
  }
}
