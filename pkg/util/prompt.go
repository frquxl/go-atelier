package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Confirm prompts the user with a message and waits for a 'yes' or 'no' answer.
// It returns true if the user confirms, false otherwise.
func Confirm(message string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [yes/no]: ", message)
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "yes" || input == "y" {
			return true
		} else if input == "no" || input == "n" {
			return false
		} else {
			fmt.Println("Please answer 'yes' or 'no'.")
		}
	}
}

// Prompt prompts the user with a message and returns their input as a string.
func Prompt(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", message)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
