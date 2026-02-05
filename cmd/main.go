package main

import (
	"os"

	"github.com/dustinbracy/learnGoWithTests/mocking"
)

func main() {
	sleeper := &mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}
