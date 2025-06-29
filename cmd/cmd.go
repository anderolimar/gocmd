package cmd

import (
	"context"
	"os"
)

type Command struct {
	Name        string
	Description string
	Usage       string
	Handler     func(ctx context.Context)
}

type Cmd interface {
	RegisterCommand(cmd Command)
	Execute()
}

func Execute() {
	cmdInstance.Execute()
}

func RegisterCommand(cmd Command) {
	cmdInstance.RegisterCommand(cmd)
}

type cmd struct {
	Commands []Command
}

var cmdInstance Cmd = &cmd{}

func (c *cmd) RegisterCommand(cmd Command) {
	c.Commands = append(c.Commands, cmd)
}

func (c *cmd) Execute() {
	args := os.Args[1:]
	if len(args) < 1 {
		c.showUsage()
		return
	}

	name := args[0]

	hasName := false
	for _, cmd := range c.Commands {
		if cmd.Name == name {
			cmd.Handler(context.Background())
			hasName = true
		}
	}
	if !hasName {
		println("Invalid command: " + name)
	}
}

func (c *cmd) showUsage() {
	println("Available commands:")
	for _, cmd := range c.Commands {
		println("	" + cmd.Name + ": " + cmd.Description)
		println("	Usage: " + cmd.Usage)
		println()
	}
}
