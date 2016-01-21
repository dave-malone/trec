package main

import "github.com/dave-malone/trec/services"

func main() {
	m := trec.NewServer()
	m.Run()
}
