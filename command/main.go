package main

import "fmt"

type Command interface {
	Execute() string
}

type Fire struct{}

func (f *Fire) Execute() string {
	return "pew pew"
}

type Jump struct{}

func (j *Jump) Execute() string {
	return "JUUUUUUMP"
}

func InputHandler(input string) string {
	commands := map[string]Command{
		"fire": &Fire{},
		"jump": &Jump{},
	}
	command := commands[input]
	if command == nil {
		return "command not found"
	}
	return command.Execute()
}

func main() {
	fmt.Printf("Fire: %s\n", InputHandler("fire"))
	fmt.Printf("Jump: %s\n", InputHandler("jump"))
	fmt.Printf("Something crazy: %s\n", InputHandler("something crazy"))
}
