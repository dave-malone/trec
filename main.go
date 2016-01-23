package main

import (
	"fmt"
	"os"

	"github.com/dave-malone/trec/services"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	fmt.Printf("Running server on port %v\n", port)

	m := trec.NewServer()
	m.RunOnAddr(":" + port)
}
