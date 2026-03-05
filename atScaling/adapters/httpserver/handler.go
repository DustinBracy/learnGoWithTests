package httpserver

import (
	"fmt"
	"net/http"

	"github.com/dustinbracy/learnGoWithTests/atScaling/domain/interactions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactions.Greet(name))
}
