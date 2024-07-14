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

var rounds = []Round{
  Round{
    Id: 1,
    Sequence: 1,
    NumCards: 10,
  },
}

var currentRoundId int = 1;

var playerRounds = []PlayerRound{
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
}

func GetAllPlayerRounds() []PlayerRound {
  return playerRounds;
}

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

// Mutations
func AddRound(numCards int) Round {
  // TODO: validate that round doesn't exist
  // TODO: track list of plauers
  currentRoundId = currentRoundId + 1;

  var newRound = Round{
    Id: currentRoundId,
    Sequence: currentRoundId,
    NumCards: numCards,
  }

  rounds = append(rounds, newRound)

  // TODO for player in player list
  playerRounds = append(playerRounds, PlayerRound{
    //Id: "1",
    Round: currentRoundId,
    Player: "Scott",
  })

  return newRound
}
