package main // tutorial has this as data which I think is a mistake

import (
  "github.com/graphql-go/graphql"
)

var playerRoundType *graphql.Object
var roundType *graphql.Object
var queryType *graphql.Object
var mutationType *graphql.Object
var Schema graphql.Schema

func init() {
  playerRoundType = graphql.NewObject(graphql.ObjectConfig{
    Name: "PlayerRound",
    Fields: graphql.Fields{
      "id": &graphql.Field{ Type: graphql.Int },
      "player": &graphql.Field{ Type: graphql.String }, // should be !
      "round": &graphql.Field{ Type: graphql.Int }, // should be !
      "bid": &graphql.Field{ Type: graphql.Int },
      "tricks": &graphql.Field{ Type: graphql.Int },
    },
  })

  roundType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Round",
    Fields: graphql.Fields{
      "id": &graphql.Field{ Type: graphql.Int },
      "sequence": &graphql.Field{ Type: graphql.Int },
      "numCards": &graphql.Field{ Type: graphql.Int },
      "playerRounds": &graphql.Field{
        Type: graphql.NewList(playerRoundType),
        Description: "All of the individual player's tricks taken and bids",
        Resolve: func(p graphql.ResolveParams) (interface{}, error){
          if round, ok := p.Source.(Round); ok { // TODO what is the point of this
            return GetPlayerRounds(round.Id), nil
          }
          return []interface{}{}, nil
        },
      },
    },
  })

  // Stub Query structure until I build the proper one
  queryType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
      "allPlayerRounds": &graphql.Field{ // deprecate?
        Type: graphql.NewList(playerRoundType),
        Description: "Gets all rounds",
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          return GetAllPlayerRounds(), nil
        },
      },
      "allRounds": &graphql.Field{
        Type: graphql.NewList(roundType),
        Description: "Gets all rounds", // TODO: later make this for a given Game
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          return GetAllRounds(), nil
        },
      },
    },
  })

  mutationType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Mutation",
    Fields: graphql.Fields{
      "createRound": &graphql.Field{
        Type: roundType,
        Description: "Create a new player-round (or, TODO, set of them)",
        Args: graphql.FieldConfigArgument{
          "numCards": &graphql.ArgumentConfig{
            Type: graphql.Int,
          },
        },
        Resolve: func(params graphql.ResolveParams) (interface{}, error) {
          return AddRound(params.Args["numCards"].(int)), nil
        },
      },
    },
  })

  Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: queryType,
    Mutation: mutationType,
  })
}
