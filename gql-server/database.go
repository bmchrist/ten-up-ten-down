package main

import (
  "math/rand/v2"
  //"fmt"
)
// put this into its own package later

type PlayerRound struct {
  Id      string  `json:"id"` // not really needed
  Round   int     `json:"text"` // points to round ID
  Player  string  `json:"text"`
  Bid     string  `json:"text"` // TODO replace with int
  Score   string  `json:"text"`
}

type Round struct {
  Id        int     `json:"id"` // not really needed - for now will just mirror the seqence number
  Sequence  int     `json:"text"`
  NumCards  string  `json:"text"` // eg 1, 10, etc
  Dealer    string  `json:"text"` // Player name/id
}

var rounds = []Round{
  Round{
    Id: 1,
    Sequence: 1,
    NumCards: "10",
  },
}

var playerRounds = []PlayerRound{
  PlayerRound{
    Id: "1",
    Round: 1,
    Player: "Ben",
    Bid: "2",
    Score: "1",
  },
  PlayerRound{
    Id: "2",
    Round: 1,
    Player: "Scott",
    Bid: "5",
    Score: "0",
  },
}

// TODO should this return Rounds that get filled with data, or the playerrounds?
func GetAllRounds() []PlayerRound {
  return playerRounds;
}

// Temporary, for mocking out
func GetRandomRound() PlayerRound {
  // todo have it grab an entire round
  return playerRounds[rand.IntN(len(playerRounds))];
}

// gets the next round ID/sequence based on the highest sequence already present
func getNextRoundId() int {
  return 1 // TODO
}

func AddRound(round string) []PlayerRound {
  // TODO: validate that round doesn't exist
  // TODO: track list of plauers
  var nextRoundId = getNextRoundId();

  rounds = append(rounds, Round{
    Id: nextRoundId,
    Sequence: nextRoundId,
    NumCards: "0", // TODO from params
  })

  // TODO for player in player list
  playerRounds = append(playerRounds, PlayerRound{
    Id: "1",
    Round: nextRoundId,
    Player: "Scott",
  })

  // should it return just the rounds that were added? or all-all
  return playerRounds
}
