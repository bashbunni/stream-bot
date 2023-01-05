package main

/* What should we want to access?
- get all commands
- retrieve a value for a command
- set a command
- edit a command
- delete a command
*/

type Repository interface {
	GetAll() []string
	GetValue(string) string
	SetValue(string)
	EditValue(string)
	DeleteValue() bool // might change this
}

type CommandsRepository struct{}

func (c *CommandsRepository) GetAll() []string       {}
func (c *CommandsRepository) GetValue(string) string {}
func (c *CommandsRepository) SetValue(string)        {}
func (c *CommandsRepository) EditValue(string)       {}
func (c *CommandsRepository) DeleteValue() bool      {}
