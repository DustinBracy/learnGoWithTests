package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	gracefulshutdown "github.com/dustinbracy/learnGoWithTests/atIntro"
)

func main() {
	var (
		ctx        = context.Background()
		httpServer = &http.Server{Addr: ":8080", Handler: http.HandlerFunc(SlowHandler)}
		server     = gracefulshutdown.NewServer(httpServer)
	)

	if err := server.ListenAndServe(ctx); err != nil {
		log.Fatalf("uh oh, didn't shut down gracefully, some responses may have been lost")
	}

	// if we get here, we know the server shut down gracefully, so we can log that
	log.Printf("server shut down gracefully, all responses were sent")
}

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Hello, world")
}
