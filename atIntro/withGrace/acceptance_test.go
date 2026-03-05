package main

import (
	"testing"
	"time"

	"github.com/dustinbracy/learnGoWithTests/atIntro/assert"
)

const (
	port = "8080"
	url  = "http://localhost:" + port
)

func TestGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := LaunchTestProgram(port)
	if err != nil {
		t.Fatalf("could not launch test program: %s", err)
	}
	t.Cleanup(cleanup)

	assert.CanGet(t, url)

	time.AfterFunc(50*time.Millisecond, func() {
		assert.NoError(t, sendInterrupt())
	})

	assert.CanGet(t, url)

	assert.CantGet(t, url)
}
