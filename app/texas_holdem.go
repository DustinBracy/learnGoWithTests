package poker

import (
	"io"
	"os"
	"time"
)

type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewTexasHoldem(store PlayerStore, alerter BlindAlerter) *TexasHoldem {
	return &TexasHoldem{
		store:   store,
		alerter: alerter,
	}
}

func (g *TexasHoldem) Start(numberOfPlayers int, to io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind, os.Stdout)
		blindTime += blindIncrement
	}
}

func (g *TexasHoldem) Finish(winner string) {
	g.store.RecordWin(winner)
}
