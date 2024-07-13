Run GraphQL server: go run main.go
Run frontend - yarn start in client-app subdir

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
