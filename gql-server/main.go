package main

import (
  "errors"
  "fmt"
  "net/http"
  "io" // used for my temporary root page, remove later
  "os"
  "github.com/graphql-go/handler"
)

// This is where we'll serve the main application from, for now it's a stub
func getRoot(w http.ResponseWriter, r *http.Request) {
  // TODO have this serve up a little html/JS snippet to make playing around
  io.WriteString(w, "This is my website!\n")
}

func main() {
  graphqlHandler := handler.New(&handler.Config{
    Schema: &Schema,
    Pretty: true,
    Playground: true,
  })

  http.HandleFunc("/", getRoot)
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
