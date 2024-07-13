package main
// put this into its own package later

type PlayerRound struct {
  Id      string  `json:"id"` // not really needed
  Player  string  `json:"text"` // replace with an ID later
  Bid     string  `json:"text"`
  Score   string  `json:"text"`
}

// Mock data
var round1 = &PlayerRound{
  Id: "1",
  Player: "Ben",
  Bid: "2",
  Score: "1",
}
var round2 = &PlayerRound{
  Id: "2",
  Player: "Scott",
  Bid: "5",
  Score: "0",
}

func GetAllRounds() []PlayerRound {
  return []PlayerRound{*round1, *round2}
}

// Temporary, for mocking out
func GetRandomRound() *PlayerRound {
  return round1
}
