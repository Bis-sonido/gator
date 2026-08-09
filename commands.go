package main

import (
	"fmt"
)

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if handler, exists := c.registeredCommands[cmd.name]; exists {
		return handler(s, cmd)
	}
	return fmt.Errorf("command not found: %s", cmd.name)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
