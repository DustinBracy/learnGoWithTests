package specifications

import (
	"testing"

	"github.com/dustinbracy/learnGoWithTests/atIntro/assert"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func GreetSpecification(t testing.TB, driver Greeter) {
	got, err := driver.Greet("Dustin")
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello, Dustin")

}
