package poker_test

import (
	"fmt"
	"testing"
	"time"

	poker "github.com/dustinbracy/learnGoWithTests/app"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts for 5 players on game start", func(t *testing.T) {

		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(dummyPlayerStore, blindAlerter)
		game.Start(5)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		checkSchedulingCases(t, blindAlerter, cases)
	})

	t.Run("schedules alerts for 7 players on game start", func(t *testing.T) {

		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(dummyPlayerStore, blindAlerter)
		game.Start(7)
		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}
		checkSchedulingCases(t, blindAlerter, cases)
	})

}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewTexasHoldem(store, dummyBlindAlerter)

	game.Finish("Ruth")
	poker.AssertPlayerWin(t, store, "Ruth")
}

func checkSchedulingCases(t *testing.T, blindAlerter *poker.SpyBlindAlerter, cases []poker.ScheduledAlert) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
			}

			got := blindAlerter.Alerts[i]
			poker.AssertScheduledAlert(t, got, want)

		})
	}
}
