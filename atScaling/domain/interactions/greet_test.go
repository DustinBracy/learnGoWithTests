package interactions_test

import (
	"testing"

	"github.com/dustinbracy/learnGoWithTests/atScaling/domain/interactions"
	"github.com/dustinbracy/learnGoWithTests/atScaling/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(interactions.Greet))
}
