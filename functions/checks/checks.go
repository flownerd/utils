package checks

import (
	"log"
	"os"
	"reflect"
	"strings"
)

// Function to check if there is error and exit from the app
func Error(err error) {
	// Check if the error is different from nil
	if err != nil {
		log.Fatalln(err)
	}
}

// Check if the interface is equal to zero
func IsZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

// Function to clean a string removing spaces, tabs and newlines
func ClearStr(s string) string {
	// Remove spaces
	clean := strings.ReplaceAll(s, " ", "")
	// Remove tabs
	clean = strings.Replace(clean, "\t", "", -1)
	// Remove newlines *Nix like
	clean = strings.Replace(clean, "\n", "", -1)
	// Remove newlines from Windows files format
	clean = strings.Replace(clean, "\r\n", "", -1)
	// Return the clean string
	return clean
}

/* FILE FUNCTIONS */

// Function to validate a file
func ObjUsable(object string) bool {

	// Variable to hold the bool of file exist
	var Exists bool

	// Check if the files exist
	if _, err := os.Stat(object); os.IsNotExist(err) {
		// The file does not exist
		Exists = false
	} else {
		// The file exist
		Exists = true
	}

	// Return the *os.File
	return Exists
}

// Function to check if the object is a directory or not
func IsDir(object string) bool {

	// Variable to return
	var isDir bool

	// Check if the object Exists if it does not exist just return false
	if !ObjUsable(object) {
		return false
	}

	// Get the Object information
	info, err := os.Stat(object)

	// Check if there is any errors
	Error(err)

	// Check if the object is a directory or not
	if info.IsDir() {
		isDir = true
	} else {
		isDir = false
	}

	// Return if is a directory or not
	return isDir
}

// Function to check if the object is a file or not
func IsFile(object string) bool {
	// Variable to return
	var isFile bool

	// Check if the object Exists if it does not exist just return false
	if !ObjUsable(object) {
		return false
	}

	// Get the Object information
	info, err := os.Stat(object)

	// Check if there is any errors
	Error(err)

	// Check if the object is a File or not
	if info.Mode().IsRegular() {
		isFile = true
	} else {
		isFile = false
	}

	// Return if is a file or not
	return isFile
}

// Function to check if the object is a symbolic links or not
func IsSymlink(object string) bool {
	// Variable to return
	var isSymlink bool

	// Check if the object Exists if it does not exist just return false
	if !ObjUsable(object) {
		return false
	}

	// Get the Object information
	info, err := os.Lstat(object)

	// Check if there is any errors
	Error(err)

	// Check if the object is a symlink or not
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		isSymlink = true
	} else {
		isSymlink = false
	}

	// Return if is a symlink or not
	return isSymlink
}

// Function to check if the parameter is empty or not.
func Empty(toCheck interface{}) bool {
	// Check if the parameter is empty or zero in the case of number
	if toCheck == "" || IsZeroOfUnderlyingType(toCheck) {
		return true
	} else {
		return false
	}
}
