package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commandRegistry := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		var words []string
		if scanner.Scan() {
			words = cleanInput(scanner.Text())
			if len(words) == 0 {
				continue
			}
		}

		userInput := words[0]

		switch userInput {
		case "exit":
			commandRegistry["exit"].callback()
		case "help":
			commandRegistry["help"].callback()
			for key, command := range commandRegistry {
				helpOutput := fmt.Sprintf("%s: %s", key, command.description)
				fmt.Println(helpOutput)
			}
		default:
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	trimmedText := strings.ToLower(strings.TrimSpace(text))

	if len(trimmedText) == 0 {
		return []string{}
	}

	words := strings.Split(trimmedText, " ")
	return words
}
