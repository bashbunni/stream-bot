package main

import (
	"reflect"
	"testing"

	"github.com/dgraph-io/badger/v3"
)

func initializeDatabaseStartingValues(t testing.TB, cmds []Command, db badger.DB, cmdRepo CommandsRepository) {
	t.Helper()
	for _, cmd := range cmds {
		if err := cmdRepo.SetValue(cmd.key, cmd.value); err != nil {
			t.Fatalf("unable to set value: %v", err)
		}
	}
}

func initializeTestingCommandRepo(t *testing.T) CommandsRepository {
	t.Helper()
	db, err := initInMemoryDB()
	if err != nil {
		t.Fatalf("unable to init in-memory DB: %v", err)
	}
	return CommandsRepository{db}
}

func TestGetAll(t *testing.T) {
	tests := []struct {
		input []Command
	}{
		{
			input: []Command{
				{[]byte("hello"), []byte("world")},
				{[]byte("milk"), []byte("cereal")},
			},
		},
	}

	for _, tc := range tests {
		c := initializeTestingCommandRepo(t)
		defer c.db.Close()

		initializeDatabaseStartingValues(t, tc.input, *c.db, c)
		got, err := c.GetAll()
		if err != nil {
			t.Fatalf("unable to get commands: %v", err)
		}
		if !reflect.DeepEqual(got, tc.input) {
			t.Fatalf("got != want;\n%v != %v", got, tc.input)
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
		c := initializeTestingCommandRepo(t)
		defer c.db.Close()

		initializeDatabaseStartingValues(t, tc.input, *c.db, c)

		got, err := c.GetValue(tc.key)
		if err != nil {
			t.Fatalf("unable to get commands: %v", err)
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got != want;\n%v != %v", got, tc.want)
		}
	}
}

func TestSetValue(t *testing.T) {
	tests := []struct {
		input    []Command
		key      []byte
		newValue []byte
	}{
		{
			input: []Command{
				{[]byte("hello"), []byte("world")},
				{[]byte("milk"), []byte("cereal")},
			},
			key:      []byte("hello"),
			newValue: []byte("chat"),
		},
		{
			input: []Command{
				{[]byte("hello"), []byte("world")},
				{[]byte("milk"), []byte("cereal")},
			},
			key:      []byte("new"),
			newValue: []byte("value"),
		},
	}

	for _, tc := range tests {
		c := initializeTestingCommandRepo(t)
		defer c.db.Close()

		initializeDatabaseStartingValues(t, tc.input, *c.db, c)

		if err := c.SetValue(tc.key, tc.newValue); err != nil {
			t.Fatalf("unable to set value: %v", err)
		}
		got, err := c.GetValue(tc.key)
		if err != nil {
			t.Fatalf("unable to get commands: %v", err)
		}
		if !reflect.DeepEqual(got, tc.newValue) {
			t.Fatalf("got != want;\n%v != %v", got, tc.newValue)
		}
	}
}

func TestDeleteValue(t *testing.T) {
	tests := []struct {
		input []Command
		key   []byte
	}{
		{
			input: []Command{
				{[]byte("hello"), []byte("world")},
				{[]byte("milk"), []byte("cereal")},
			},
			key: []byte("hello"),
		},
		{
			input: []Command{
				{[]byte("hello"), []byte("world")},
				{[]byte("milk"), []byte("cereal")},
			},
			key: []byte("new"), // TODO: why does this pass
		},
	}

	for _, tc := range tests {
		c := initializeTestingCommandRepo(t)
		defer c.db.Close()

		initializeDatabaseStartingValues(t, tc.input, *c.db, c)

		if err := c.DeleteValue(tc.key); err != nil {
			t.Fatalf("unable to delete value: %v", err)
		}
		if v, _ := c.GetValue(tc.key); string(v) != "this command does not exist" {
			t.Fatalf("command still exists after being deleted")
		}
	}
}

func initInMemoryDB() (*badger.DB, error) {
	opt := badger.DefaultOptions("").WithInMemory(true).WithLogger(nil)
	return badger.Open(opt)
}
