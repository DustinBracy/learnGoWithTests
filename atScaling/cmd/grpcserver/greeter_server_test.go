package main_test

import (
	"fmt"
	"testing"

	"github.com/dustinbracy/learnGoWithTests/atScaling/adapters"
	"github.com/dustinbracy/learnGoWithTests/atScaling/adapters/grpcserver"
	"github.com/dustinbracy/learnGoWithTests/atScaling/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	var (
		port   = "50051"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)
	t.Cleanup(driver.Close)
	adapters.StartDockerServer(t, port, "grpcserver")
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
