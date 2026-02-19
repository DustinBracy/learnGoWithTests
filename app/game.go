package poker

// Game manages the state of a poker game, including starting the game and recording wins.
type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}
