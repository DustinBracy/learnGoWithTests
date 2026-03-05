package generics_test

import (
	"testing"

	"github.com/dustinbracy/learnGoWithTests/generics"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on ints", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "world")
	})

	// AssertEqual(t, 1, "1") // uncomment to see error

}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v want false", got)
	}
}

func TestStack(t *testing.T) {
	t.Run("int stack", func(t *testing.T) {
		intStack := generics.NewStack[int]()

		// check empty
		AssertTrue(t, intStack.IsEmpty())

		// add something, check not empty
		intStack.Push(123)
		AssertFalse(t, intStack.IsEmpty())

		// add another, pop and check value
		intStack.Push(456)
		value, _ := intStack.Pop()
		AssertEqual(t, value, 456)

		value, _ = intStack.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, intStack.IsEmpty())

		// can get numbers back as numbers
		intStack.Push(1)
		intStack.Push(2)
		num1, _ := intStack.Pop()
		num2, _ := intStack.Pop()
		AssertEqual(t, num1+num2, 3)
	})
}
