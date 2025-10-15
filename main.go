package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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

		command, exists := getCommandRegistry()[userInput]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
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
