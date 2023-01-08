package main

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

func main() {
	// do nothing. Eventually start the bots
	opts := badger.DefaultOptions("/tmp/badger")
	opts.Logger = nil
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	c := CommandsRepository{db}
	if err := c.SetValue([]byte("hello"), []byte("world")); err != nil {
		log.Fatal(err)
	}
	cmds, err := c.GetAll()
	for _, val := range cmds {
		fmt.Printf("key: %s, val: %s\n", val.key, val.value)
	}
	if err != nil {
		log.Fatal(err)
	}
}
