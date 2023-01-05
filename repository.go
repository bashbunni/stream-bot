package main

import (
	"fmt"

	"github.com/dgraph-io/badger/v3"
)

/* What should we want to access?
- get all commands
- retrieve a value for a command
- set a command
- edit a command
- delete a command
*/

type Command struct {
	key   []byte
	value []byte
}

type Repository interface {
	GetAll() []string
	GetValue(string) string
	SetValue(string)
	EditValue(string)
	DeleteValue() bool // might change this
}

type CommandsRepository struct {
	db *badger.DB
}

func (c *CommandsRepository) GetAll() ([]Command, error) {
	var cmds []Command
	err := c.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v) // TODO: remove this
				cmds = append(cmds, Command{key: k, value: v})
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return cmds, nil
}

/*
func (c *CommandsRepository) GetValue(string) string {}
func (c *CommandsRepository) SetValue(string)        {}
func (c *CommandsRepository) EditValue(string)       {}
func (c *CommandsRepository) DeleteValue() bool      {}
*/
