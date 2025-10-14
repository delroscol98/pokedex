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
		var userInput string
		if scanner.Scan() {
			userInput = scanner.Text()
		}
		userWords := cleanInput(userInput)

		fmt.Println("Your command was:", userWords[0])
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
