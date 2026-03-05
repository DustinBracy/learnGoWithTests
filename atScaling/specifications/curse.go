package specifications

import (
	"testing"

	"github.com/dustinbracy/learnGoWithTests/atIntro/assert"
)

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t testing.TB, meany MeanGreeter) {
	got, err := meany.Curse("Dustin")
	assert.NoError(t, err)
	assert.Equal(t, got, "Go away, Dustin")
}
