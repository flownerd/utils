package inputs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/flownerd/keyboard"
	"github.com/flownerd/utils/functions/checks"
)

// Global vars
var (
	reader *bufio.Reader
)

// Function to show a simple prompt
func prompt() {
	fmt.Print("-> ")
}

// Function to read the user input and return a string
func ReadString(format string, a ...interface{}) string {

	// Run until get a value different from empty
	for {

		// Show the message sent by the user
		fmt.Printf(format, a...)
		prompt()

		// Get the input clean
		userInput := clearInput()

		// Check if the input is empty or not
		if userInput != "" {
			return userInput
		}
	}
}

// Function to read the user input and return a integer
func ReadInt(format string, a ...interface{}) int {
	// Run until get a value different from empty
	for {

		// Show the message sent by the user
		fmt.Printf(format, a...)
		prompt()

		// Get the input clean
		userInput := clearInput()

		// Convert the input to int
		num, err := strconv.Atoi(userInput)

		// Check if the input is empty or not
		if err == nil {
			return num
		}
	}
}

// Function to read the user input and return a float64
func ReadFloat(format string, a ...interface{}) float64 {
	// Run until get a value different from empty
	for {

		// Show the message sent by the user
		fmt.Printf(format, a...)
		prompt()

		// Get the input clean
		userInput := clearInput()

		// Convert the input to int
		num, err := strconv.ParseFloat(userInput, 64)

		// Check if the input is empty or not
		if err == nil {
			return num
		}
	}
}

// Function to read the user input and return a bool
func ReadBool(format string, a ...interface{}) bool {

	// Open the keyboard to be handle
	err := keyboard.Open()
	checks.Error(err)

	// Defer the keyboard
	defer keyboard.Close()

	// Run until get a value different from empty
	for {

		// Show the message sent by the user
		fmt.Printf(format, a...)
		prompt()

		// Here we need to read a single key
		char, _, err := keyboard.GetSingleKey()
		checks.Error(err)

		// Check if we have the y or n
		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please type y or n")
		} else if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		}

	}
}

// Function to clean the userInput
func clearInput() string {
	// Calling the Reader
	reader = bufio.NewReader(os.Stdin)

	// Read the input and clean it up
	userInput, _ := reader.ReadString('\n')
	userInput = strings.ReplaceAll(userInput, "\r\n", "")
	userInput = strings.ReplaceAll(userInput, "\n", "")

	// return the input without the new line
	return userInput
}
