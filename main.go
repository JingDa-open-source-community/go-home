package main

import (
	"os"

	"github.com/go-home/call"
)

func main() {
	commands := call.Commands
	args := os.Args

	for _, cmd := range commands {
		if cmd.Run != nil && cmd.Name() == args[1] {
			cmd.Flag.Parse(args[2:])
			args = cmd.Flag.Args()
			os.Exit(cmd.Run(cmd, args))
		}
	}
}
