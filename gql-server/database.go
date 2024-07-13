package main
// put this into its own package later

type PlayerRound struct {
  Id      string  `json:"id"` // not really needed
  Player  string  `json:"text"` // replace with an ID later
  Bid     string  `json:"text"`
  Score   string  `json:"text"`
}

// Mock data
var testRound = &PlayerRound{
  Id: "1",
  Player: "Ben",
  Bid: "2",
  Score: "1",
}

// Temporary, for mocking out
func RandomRound() *PlayerRound {
  return testRound
}
