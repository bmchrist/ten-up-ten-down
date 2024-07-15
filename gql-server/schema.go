package main // tutorial has this as data which I think is a mistake

import (
  "github.com/graphql-go/graphql"
  "fmt"
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
      "id": &graphql.Field{
        Type: graphql.String,

			  // We want a unique ID for react to understand when it was updated
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          if playerRound, ok := p.Source.(PlayerRound); ok {
						return fmt.Sprintf("%s-%d", playerRound.Player, playerRound.Round), nil
					}
					return nil, nil
        },
      },
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

  queryType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
      //"allPlayerRounds": &graphql.Field{ // deprecate?
        //Type: graphql.NewList(playerRoundType),
        //Description: "Gets all rounds",
        //Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          //return GetAllPlayerRounds(), nil
        //},
      //},
      "allRounds": &graphql.Field{
        Type: graphql.NewList(roundType),
        Description: "Gets all rounds", // TODO: later make this for a given Game
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          return GetAllRounds(), nil
        },
      },
    },
  })

  // Adds a new round, incrementing the sequence/id, with a provided number of cards
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
      "updateBid": &graphql.Field{
        Type: playerRoundType,
        Description: "Sets the bid for a given player",
        Args: graphql.FieldConfigArgument{
          "roundId": &graphql.ArgumentConfig{
            Type: graphql.Int,
          },
          "player": &graphql.ArgumentConfig{
            Type: graphql.String,
          },
          "bid": &graphql.ArgumentConfig{
            Type: graphql.Int,
          },
        },
        Resolve: func(params graphql.ResolveParams) (interface{}, error) {
          return UpdateBid(
            params.Args["roundId"].(int),
            params.Args["player"].(string),
            params.Args["bid"].(int),
          ), nil
        },
      },
      "updateTricks": &graphql.Field{
        Type: playerRoundType,
        Description: "Sets the tricks scored for a given player",
        Args: graphql.FieldConfigArgument{
          "roundId": &graphql.ArgumentConfig{
            Type: graphql.Int,
          },
          "player": &graphql.ArgumentConfig{
            Type: graphql.String,
          },
          "tricks": &graphql.ArgumentConfig{
            Type: graphql.Int,
          },
        },
        Resolve: func(params graphql.ResolveParams) (interface{}, error) {
          return UpdateTricks(
            params.Args["roundId"].(int),
            params.Args["player"].(string),
            params.Args["tricks"].(int),
          ), nil
        },
      },
    },
  })

  Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: queryType,
    Mutation: mutationType,
  })
}
