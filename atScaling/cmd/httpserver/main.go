package main

import (
	"log"
	"net/http"

	"github.com/dustinbracy/learnGoWithTests/atScaling/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
