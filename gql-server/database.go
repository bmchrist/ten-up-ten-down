package main

import ("math/rand/v2")
// put this into its own package later

type PlayerRound struct {
  Id      string  `json:"id"` // not really needed
  Player  string  `json:"text"` // replace with an ID later
  Bid     string  `json:"text"`
  Score   string  `json:"text"`
}

var rounds = []PlayerRound{
  PlayerRound{
    Id: "1",
    Player: "Ben",
    Bid: "2",
    Score: "1",
  },
  PlayerRound{
    Id: "2",
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
