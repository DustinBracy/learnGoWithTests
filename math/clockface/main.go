package main

import (
	"os"
	"time"

	"github.com/dustinbracy/learnGoWithTests/math/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
