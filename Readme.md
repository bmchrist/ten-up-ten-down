Run GraphQL server: `go run .` in gql-server

Run frontend - `yarn start` in ten-up-client-app subdir

A game has players and rounds
In a round each player has a bid and a score

Game
- Id
- DateTime
- Has many PlayerRounds
- Has many Players (not re-used outside of games)

PlayerRounds
- ID: Round+Player
- Round (-10 to -1, 1 to 10)
- Bid
- Score
- Player (string name as the ID, can later be joined to some table)
- Game


TODO
- Render them
- Have proper ID structure for PlayerRound
- Should we build all rounds at game start, or build them as we go?
 - Probably add as we go
- Allow mutating a given playerround's score and bid

- Switch to a persistent data store
- Add "Game" scoping (instead of player rounds being global)


Ideas
- Button to hit when player takes a trick
-



Frontend should
Mutations:
- Create Game
- Add players
- Add round (uses current player list, sets dealer based on prev round, take # of cards for round)
- Add Score for given player+round
- Add Bid for given player+round

Fetching
- Get all rounds and round info and display, one query
- - at first have this just automatically get subqueries using gql, but look into how we don't n+
- Know who the dealer is for a given round


TBD
- Frontend or backend track total score


Add tests for the go server functions
and then switch to a new data structure
