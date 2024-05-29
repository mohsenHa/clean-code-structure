package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type command struct {
	description string
	function    func()
}

var commands = map[int]command{}

func init() {
	i := 1

	commands[i] = command{
		description: "Create a all",
		function:    newAll,
	}
	i++

	commands[i] = command{
		description: "Create handler",
		function:    handler,
	}
	i++

	commands[i] = command{
		description: "Create param",
		function:    param,
	}
	i++

	commands[i] = command{
		description: "Create service",
		function:    service,
	}
	i++

	commands[i] = command{
		description: "Create validator",
		function:    validator,
	}
	i++

	commands[i] = command{
		description: "Exit",
		function: func() {
			fmt.Println("See you later")
			os.Exit(0)
		},
	}
}

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("What do you want to do?")
		for i := 1; i < len(commands)+1; i++ {
			c := commands[i]
			fmt.Printf("%d. %s\n", i, c.description)
		}
		scanner.Scan()
		commandId := scanner.Text()
		atoi, err := strconv.Atoi(commandId)
		if err != nil {
			fmt.Println(err)

			continue
		}
		c, ok := commands[atoi]
		if !ok {
			fmt.Println("Unknown command!")
			continue
		}
		c.function()
	}

}
