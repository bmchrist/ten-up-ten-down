package main

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

//var roundsStore = []Round {
//}
//var playerRoundsStore = map[Round][]PlayerRound {
//}

func init() {
  SeedRounds();
}

func SeedRounds() bool {
  AddPlayer("Ben")
  AddPlayer("Scott")

  var round = AddRound(10)
  UpdateBid(round.Id, "Ben", 1)
  UpdateTricks(round.Id, "Ben", 0)
  UpdateBid(round.Id, "Scott", 2)
  UpdateTricks(round.Id, "Scott", 2)

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
// Add player, returns list of all players
// TODO expose this via gql
func AddPlayer(name string) []string {
  players = append(players, name)
  return players
}

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

// TODO abstract out the finding function, use pointer for both of the following
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

// How should I return a falied status? For now just returns a null PlayerRound
func UpdateTricks(roundId int, player string, tricks int) PlayerRound {
  var updatedPlayerRound PlayerRound
  for i:=0; i < len(playerRounds); i++ { // add or playerRound not null to exit if found already
    if playerRounds[i].Round == roundId && playerRounds[i].Player == player {
      playerRounds[i].Tricks = tricks
      updatedPlayerRound = playerRounds[i]
    }
  }
  return updatedPlayerRound
}
