package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	poker "github.com/dustinbracy/learnGoWithTests/app"
)

var (
	dummyBlindAlerter = &poker.SpyBlindAlerter{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
	dummyStdIn        = &bytes.Buffer{}
	dummyStdOut       = &bytes.Buffer{}
)

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
	FinishCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int, to io.Writer) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
	g.FinishCalled = true
}
func TestCLI(t *testing.T) {

	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("3", "Chris wins")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt+"\n")
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &GameSpy{}

		in := userSends("8", "Cleo wins")
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("pies")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, "\n"+poker.InvalidPlayersErrMsg)
	})
}

func assertGameStartedWith(t *testing.T, game *GameSpy, numberOfPlayers int) {
	t.Helper()

	if game.StartedWith != numberOfPlayers {
		t.Errorf("game started with %d players, want %d", game.StartedWith, numberOfPlayers)
	}
}

func assertFinishCalledWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()

	if game.FinishedWith != winner {
		t.Errorf("game finished with winner %q, want %q", game.FinishedWith, winner)
	}
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()

	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout, but expected %q", got, want)
	}
}
