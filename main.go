package main

import (
	"os"

	"github.com/dave-malone/trec/services"
	"github.com/xchapter7x/lo"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	lo.G.Debug("Running server on port %v\n", port)

	m := trec.NewServer()
	m.RunOnAddr(":" + port)
}
