package inputs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Input struct that will hold the default buffer and the option to show or not the output of the user message.
type Input struct {
	Buffer        *bufio.Reader // Define the buffer that will be used throught the package
	NotShowOutput bool          // Define the variable that will enable or disable the output of the message
}

// Function to show a simple prompt (-> )
func prompt() {
	fmt.Print("-> ")
}

// Receiver to define the buffer that will be used throught the package
func (i *Input) New(reader *bufio.Reader) *Input {
	// Set the reader
	if reader == nil {
		// If the reader was not passed use the stdin
		i.Buffer = bufio.NewReader(os.Stdin)
	} else {
		i.Buffer = reader
	}

	// Return the pointer of the struct
	return i
}

// Function to read Data from the user
func (i *Input) readData(format string, a ...interface{}) string {

	// Show the message sent by the user
	if !i.NotShowOutput {
		fmt.Printf(format, a...)
		prompt()
	}

	// Get the input clean
	userInput := i.clearInput()

	// Return the user input
	return userInput
}

// Function to read the user input and return a string
func (i *Input) ReadString(format string, a ...interface{}) string {

	// Run until get a value different from empty
	for {

		// Read the data from the reader
		userInput := i.readData(format, a...)

		// Check if the input is empty or not
		if userInput != "" {
			return userInput
		}
	}
}

// Function to read the user input and return a integer
func (i *Input) ReadInt(format string, a ...interface{}) int {
	// Run until get a value different from empty
	for {

		// Read the data from the reader
		userInput := i.readData(format, a...)

		// Convert the input to int
		num, err := strconv.Atoi(userInput)

		// Check if the input is empty or not
		if err == nil {
			return num
		}
	}
}

// Function to read the user input and return a float64
func (i *Input) ReadFloat(format string, a ...interface{}) float64 {
	// Run until get a value different from empty
	for {

		// Read the data from the reader
		userInput := i.readData(format, a...)

		// Convert the input to int
		num, err := strconv.ParseFloat(userInput, 64)

		// Check if the input is empty or not
		if err == nil {
			return num
		}
	}
}

// Function to read the user input and return a bool
func (i *Input) ReadBool(format string, a ...interface{}) bool {

	// Run until get a value different from empty
	for {

		// Read the data from the reader
		userInput := i.readData(format, a...)

		// Check if we have the y or n
		if strings.ToLower(userInput) != "y" && strings.ToLower(userInput) != "n" {
			fmt.Println("Please type y or n")
		} else if userInput == "n" || userInput == "N" {
			return false
		} else if userInput == "y" || userInput == "Y" {
			return true
		}

	}
}

// Function to clear the userInput
func (i *Input) clearInput() string {

	// Read the input and clean it up
	userInput, _ := i.Buffer.ReadString('\n')
	userInput = strings.ReplaceAll(userInput, "\r\n", "")
	userInput = strings.ReplaceAll(userInput, "\n", "")

	// return the input without the new line
	return userInput
}
