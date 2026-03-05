package interactions_test

import (
	"testing"

	"github.com/dustinbracy/learnGoWithTests/atScaling/domain/interactions"
	"github.com/dustinbracy/learnGoWithTests/atScaling/specifications"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(t, specifications.CurseAdapter(interactions.Curse))
}
