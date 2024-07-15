package main

import (
  "testing"
)

// Relies on SeedData() being run
func TestGetPlayerRounds(t *testing.T) {
  // Properly gets both player's round for Round 1
  var foundPlayerRounds = GetPlayerRounds(1)
  if len(playerRounds) != 2 {
    t.Fatalf(`Player rounds expected 2, found %d`, len(foundPlayerRounds))
  }

  // Returns an empty slice if no rounds found
  foundPlayerRounds = GetPlayerRounds(2)
  if len(foundPlayerRounds) != 0 {
    t.Fatalf(`Player rounds expected 0, found %d`, len(foundPlayerRounds))
  }
}

func TestAddRound(t *testing.T){
  var allRoundsInitial = GetAllRounds()
  var newRound = AddRound(5)
  var allRoundsAfter = GetAllRounds()

  var countNewRounds = len(allRoundsAfter)-len(allRoundsInitial)
  if countNewRounds != 1 {
    t.Fatalf(`Expected one additional round, found %d`, countNewRounds)
  }

  if newRound.NumCards != 5 {
    t.Fatalf(`Expected numCards = 5 on new round, found %d`, newRound.NumCards)
  }

  // TODO Check the rounds list to see if the round exists with the values expected
}

func TestUpdateBid(t *testing.T){
  // TODO
}

func TestUpdateTricks(t *testing.T){
  // TODO
}

//TODO: test actual gql queries
