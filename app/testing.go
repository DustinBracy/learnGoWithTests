package poker

import (
	"fmt"
	"io"
	"reflect"
	"testing"
	"time"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("expected a win call but got %d", len(store.WinCalls))
	}

	if store.WinCalls[0] != winner {
		t.Errorf("didn't record correct winner, got %q want %q", store.WinCalls[0], winner)
	}
}

func AssertScheduledAlert(t *testing.T, got, want ScheduledAlert) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s *ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, ScheduledAlert{at, amount})
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
