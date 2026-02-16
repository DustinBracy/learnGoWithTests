package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/dustinbracy/learnGoWithTests/atScaling/adapters"
	"github.com/dustinbracy/learnGoWithTests/atScaling/adapters/httpserver"
	"github.com/dustinbracy/learnGoWithTests/atScaling/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	var (
		port    = "8080"
		baseURL = fmt.Sprintf("http://127.0.0.1:%s", port)
		driver  = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{Timeout: 1 * time.Second}}
	)
	adapters.StartDockerServer(t, port, "httpserver")
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
