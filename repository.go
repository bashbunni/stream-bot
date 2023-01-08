package main

import (
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
	GetAll() ([]Command, error) // TODO ask about this
	GetValue([]byte) []byte
	SetValue(key, value []byte)
	EditValue([]byte)
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

func (c *CommandsRepository) SetValue(k, v []byte) error {
	err := c.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(k, v)
		return err
	})
	return err
}

func (c *CommandsRepository) GetValue(k []byte) ([]byte, error) {
	var v []byte
	err := c.db.View(func(txn *badger.Txn) error {
		i, err := txn.Get(k)
		if err != nil {
			v = []byte("this command does not exist")
			return nil
		}
		v, err = i.ValueCopy(v)
		return err
	})
	return v, err
}

/*
func (c *CommandsRepository) DeleteValue() bool {}
*/
