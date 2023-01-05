package main

import (
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

func main() {
	// do nothing. Eventually start the bots
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
