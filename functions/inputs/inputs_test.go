package inputs

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// User struct just to work throughout the tests
type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	OwnsADog       bool
}

// Global vars
var (
	notShowOutput = true        // Variable to set if the output message for the user will be showed or not
	msg           string        // Message will be hold here
	reader        *bufio.Reader // Define the reader
)

// Function to test the function ReadString
func TestReadString(t *testing.T) {

	// Define the user struct
	var user User

	// Input that will be passed to the function
	userInput := "FlowNerd\n"

	// Set the Reader
	input := setReader(userInput)

	// Define the message
	msg = "What is your name?"

	// Read the data from the buffer
	user.UserName = input.ReadString(fmt.Sprintln(msg))

	// Check if the type returned is the correct one
	if reflect.TypeOf(user.UserName).Kind() == reflect.String {
		t.Logf("%v [%s]", msg, user.UserName)
	} else {
		t.Errorf("%v [%s]", msg, user.UserName)
	}

}

// Function to test the function ReadInt
func TestReadInt(t *testing.T) {

	// Define the user struct
	var user User

	// Input that will be passed to the function
	userInput := "33\n"

	// Set the Reader
	input := setReader(userInput)

	// Define the message
	msg = "How old are you?"

	// Read the data from the buffer
	user.Age = input.ReadInt(fmt.Sprintln(msg))

	// Check if the type returned is the correct one
	if reflect.TypeOf(user.Age).Kind() == reflect.Int {
		t.Logf("%v [%d]", msg, user.Age)
	} else {
		t.Errorf("%v [%d]", msg, user.Age)
	}

}

// Function to test the function ReadFloat
func TestReadFloat(t *testing.T) {

	// Define the user struct
	var user User

	// Input that will be passed to the function
	userInput := "3.141559\n"

	// Set the Reader
	input := setReader(userInput)

	// Define the message
	msg = "What is your favorite number?"

	// Read the data from the buffer
	user.FavoriteNumber = input.ReadFloat(fmt.Sprintln(msg))

	// Check if the type returned is the correct one
	if reflect.TypeOf(user.FavoriteNumber).Kind() == reflect.Float64 {
		t.Logf("%v [%.6f]", msg, user.FavoriteNumber)
	} else {
		t.Errorf("%v [%.6f]", msg, user.FavoriteNumber)
	}

}

// Function to test the function ReadBool
func TestReadBool(t *testing.T) {

	// Define the user struct
	var user User

	// Input that will be passed to the function
	userInput := "y\n"

	// Set the Reader
	input := setReader(userInput)

	// Define the message
	msg = "Do you own a dog (y/n)?"

	// Read the data from the buffer
	user.OwnsADog = input.ReadBool(fmt.Sprintln(msg))

	// Check if the type returned is the correct one
	if reflect.TypeOf(user.OwnsADog).Kind() == reflect.Bool {
		t.Logf("%v [%t]", msg, user.OwnsADog)
	} else {
		t.Errorf("%v [%t]", msg, user.OwnsADog)
	}

}

// Function to set the Reader
func setReader(userInput string) Input {
	// Convert the userInput to the *bufio.Reader
	reader = bufio.NewReader(strings.NewReader(userInput))

	// Create a variable from Struct Inputs
	input := Input{
		NotShowOutput: notShowOutput,
		Buffer:        reader,
	}

	// Return the object
	return input
}
