package main

//import (
  //"math/rand/v2"
  //"fmt"
//)

type PlayerRound struct {
  Id      string  `json:"id"` // not really needed
  Round   int     `json:"integer"` // points to round ID
  Player  string  `json:"text"`
  Bid     int     `json:"integer"` // TODO replace with int
  Tricks  int     `json:"integer"`
}

type Round struct {
  Id        int     `json:"id"` // not really needed - for now will just mirror the seqence number
  Sequence  int     `json:"integer"`
  NumCards  int     `json:"integer"` // eg 1, 10, etc
  Dealer    string  `json:"text"` // Player name/id
}

// Consider storing the playerRounds under rounds in a map..
var rounds = []Round{}
var playerRounds = []PlayerRound{}
var currentRoundId int = 0;
var players = []string{}

func init() {
  // For testing, add some seed data. Defined in database.go
  SeedRounds();
}

func SeedRounds() bool {
  currentRoundId = 1;

  players = []string{"Ben", "Scott"};

  rounds = append(rounds,
    Round{
      Id: 1,
      Sequence: 1,
      NumCards: 10,
    },
  )

  playerRounds = append(
    playerRounds,
    PlayerRound{
      Id: "1",
      Round: 1,
      Player: "Ben",
      Bid: 2,
      Tricks: 1,
    },
    PlayerRound{
      Id: "2",
      Round: 1,
      Player: "Scott",
      Bid: 5,
      Tricks: 0,
    },
  )

  return true
}

// -------
// Getters
// -------

// Likely not needed anymore (we get these via the Rounds which uses the GetPlayerRounds call)
//func GetAllPlayerRounds() []PlayerRound {
  //return playerRounds;
//}

// Gets all the player rounds for a given round ID (which is the same as its sequence)
func GetPlayerRounds(round_id int) []PlayerRound {
  var selectedRounds = []PlayerRound{}
  for i:=0; i < len(playerRounds); i++ {
    if playerRounds[i].Round == round_id {
      selectedRounds = append(selectedRounds, playerRounds[i])
    }
  }
  return selectedRounds;
}

func GetAllRounds() []Round {
  return rounds;
}

// ---------
// Mutations
// ---------
func AddRound(numCards int) Round {
  // TODO: validate that round doesn't exist
  currentRoundId = currentRoundId + 1;

  var newRound = Round{
    Id: currentRoundId,
    Sequence: currentRoundId,
    NumCards: numCards,
  }

  rounds = append(rounds, newRound)

  for i:=0; i < len(players); i++ {
    playerRounds = append(playerRounds, PlayerRound{
      Round: currentRoundId,
      Player: players[i],
    })
  }

  return newRound
}

// How should I return a falied status? For now just returns a null PlayerRound
func UpdateBid(roundId int, player string, bid int) PlayerRound {
  var updatedPlayerRound PlayerRound
  for i:=0; i < len(playerRounds); i++ { // add or playerRound not null to exit if found already
    if playerRounds[i].Round == roundId && playerRounds[i].Player == player {
      playerRounds[i].Bid = bid
      updatedPlayerRound = playerRounds[i]
    }
  }
  return updatedPlayerRound
}
