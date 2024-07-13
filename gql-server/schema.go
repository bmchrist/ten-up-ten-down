package main // tutorial has this as data which I think is a mistake

// TODO understand more about order that multiple files get used, pros and cons of importing in diff places, etc

import (
  "github.com/graphql-go/graphql"
)

var playerRoundType *graphql.Object
var queryType *graphql.Object
var Schema graphql.Schema

func init() {
  playerRoundType = graphql.NewObject(graphql.ObjectConfig{
    Name: "PlayerRound",
    Fields: graphql.Fields{
      "id": &graphql.Field{ Type: graphql.Int },
      "player": &graphql.Field{ Type: graphql.String }, // should be !
      "bid": &graphql.Field{ Type: graphql.Int },
      "score": &graphql.Field{ Type: graphql.Int },
    },
  })

  // Stub Query structure until I build the proper one
  queryType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
      "allRounds": &graphql.Field{
        Type: graphql.NewList(playerRoundType),
        Description: "Gets all rounds", // TODO: later make this for a given Game
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          return GetAllRounds(), nil
        },
      },
      "randomRound": &graphql.Field{
        Type: playerRoundType,
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          return GetRandomRound(), nil
        },
      },
    },
  })

  Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: queryType,
  })
}
