package main

import (
	"reflect"
	"testing"

	"github.com/dgraph-io/badger/v3"
)

func TestGetAll(t *testing.T) {
	tests := []struct {
		want []Command
	}{
		{
			want: []Command{
				{[]byte("hello"), []byte("world")},
				{[]byte("milk"), []byte("cereal")},
			},
		},
	}

	for _, tc := range tests {
		// create in-memory db for testing
		// add values
		// check that the lists are equal

		db, err := InitInMemoryDB()
		if err != nil {
			t.Fatalf("unable to init in-memory DB: %v", err)
		}
		defer db.Close()
		c := CommandsRepository{db}

		for _, cmd := range tc.want {
			if err := c.SetValue(cmd.key, cmd.value); err != nil {
				t.Fatalf("unable to set value: %v", err)
			}
		}
		got, err := c.GetAll()
		if err != nil {
			t.Fatalf("unable to get commands: %v", err)
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got != want;\n%v != %v", got, tc.want)
		}
	}
}

func InitInMemoryDB() (*badger.DB, error) {
	opt := badger.DefaultOptions("").WithInMemory(true)
	return badger.Open(opt)
}
