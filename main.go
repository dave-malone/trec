package main

import (
	"os"

	"github.com/xchapter7x/lo"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	lo.G.Debug("Running server on port %v\n", port)

	n := NewServer()
	n.Run(":3000")
}
