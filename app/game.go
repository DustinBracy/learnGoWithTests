package poker

import "io"

// Game manages the state of a poker game, including starting the game and recording wins.
type Game interface {
	Start(numberOfPlayers int, alertsDestination io.Writer)
	Finish(winner string)
}
