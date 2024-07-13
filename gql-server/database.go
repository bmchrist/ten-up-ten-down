package main

import (
  "math/rand/v2"
  //"fmt"
)
// put this into its own package later

type PlayerRound struct {
  Id      string  `json:"id"` // not really needed
  Round   string  `json:"text"`
  Player  string  `json:"text"`
  Bid     string  `json:"text"` // TODO replace with int
  Score   string  `json:"text"`
}

var rounds = []PlayerRound{
  PlayerRound{
    Id: "1",
    Round: "-10",
    Player: "Ben",
    Bid: "2",
    Score: "1",
  },
  PlayerRound{
    Id: "2",
    Round: "-10",
    Player: "Scott",
    Bid: "5",
    Score: "0",
  },
}

func GetAllRounds() []PlayerRound {
  return rounds;
}

// Temporary, for mocking out
func GetRandomRound() PlayerRound {
  return rounds[rand.IntN(len(rounds))];
}

func AddRound(round string) []PlayerRound {
  // TODO: validate that round doesn't exist
  // TODO: track list of plauers
  rounds = append(rounds, PlayerRound{ 
    Id: "1", 
    Round: round,
    Player: "Scott",
  })
  rounds = append(rounds, PlayerRound{ 
    Id: "1", 
    Round: round,
    Player: "Ben",
  })
  return rounds 
}
