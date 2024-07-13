package main

import (
  "errors"
  "fmt"
  "net/http"
  "os"
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/handler"
 )

var queryType = graphql.NewObject(graphql.ObjectConfig{
  Name: "Query",
  Fields: graphql.Fields{
    "latestPost": &graphql.Field{
      Type: graphql.String,
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        return "Hello world", nil
      },
    },
  },
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
  Query: queryType,
})

func main() {

  h := handler.New(&handler.Config{
    Schema: &Schema,
    Pretty: true,
  })

  http.Handle("/graphql", h)

  err := http.ListenAndServe(":3333", nil)

  if errors.Is(err, http.ErrServerClosed) {
    fmt.Printf("server closed\n")
  } else if err != nil {
    fmt.Printf("error starting server: %s\n", err)
    os.Exit(1)
  }
}
