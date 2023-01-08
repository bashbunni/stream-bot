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

func TestGetValue(t *testing.T) {
	tests := []struct {
		input []Command
		key   []byte
		want  []byte
	}{
		{
			input: []Command{
				{[]byte("hello"), []byte("world")},
				{[]byte("milk"), []byte("cereal")},
			},
			key:  []byte("hello"),
			want: []byte("world"),
		},
		{
			input: []Command{
				{[]byte("hello"), []byte("world")},
				{[]byte("milk"), []byte("cereal")},
			},
			key:  []byte("i don't exist"),
			want: []byte("this command does not exist"),
		},
	}

	for _, tc := range tests {
		db, err := InitInMemoryDB()
		if err != nil {
			t.Fatalf("unable to init in-memory DB: %v", err)
		}
		defer db.Close()
		c := CommandsRepository{db}

		for _, cmd := range tc.input {
			if err := c.SetValue(cmd.key, cmd.value); err != nil {
				t.Fatalf("unable to set value: %v", err)
			}
		}
		got, err := c.GetValue(tc.key)
		if err != nil {
			t.Fatalf("unable to get commands: %v", err)
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got != want;\n%v != %v", got, tc.want)
		}
	}
}

func InitInMemoryDB() (*badger.DB, error) {
	opt := badger.DefaultOptions("").WithInMemory(true).WithLogger(nil)
	return badger.Open(opt)
}
