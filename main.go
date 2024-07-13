package main

import (
  "errors"
  "fmt"
  "io"
  "net/http"
  "os"
 )

func getRoot(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("got / request\n")
  io.WriteString(w, "This is my website!\n")
}

func getSecondPage(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("got / other page request\n")
  io.WriteString(w, "This is my other page of my website!\n")
}

func main() {
  http.HandleFunc("/", getRoot)
  http.HandleFunc("/hello", getSecondPage)

  err := http.ListenAndServe(":3333", nil)

  if errors.Is(err, http.ErrServerClosed) {
    fmt.Printf("server closed\n")
  } else if err != nil {
    fmt.Printf("error starting server: %s\n", err)
    os.Exit(1)
  }
}
