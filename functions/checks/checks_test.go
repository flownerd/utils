package checks

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

var (
	// Variable to hold the output of the check
	is_ok bool

	// Message will be hold here
	msg string

	// Variable to hold the number zero inside an interface
	isZero interface{}
)

// Function to test the IsSymlink
func TestIsSymlink(t *testing.T) {
	// Synlink Name
	var symlinkname string

	// Getting the local temp dirs
	tempDir := os.TempDir()

	// Create a new temp directory
	dir, err := ioutil.TempDir(tempDir, "isfile")
	checkError(err)

	// Remove the temp dir after the function is done
	defer os.RemoveAll(dir)

	// Create a temp file inside the temp directory
	file, err := ioutil.TempFile(dir, "")
	checkError(err)

	// Creating the symlink pointing to the tempfile
	symlinkname = fmt.Sprintf("%v-%v", file.Name(), "link")
	os.Symlink(file.Name(), symlinkname)

	// Check if the object is a symlink
	is_ok = IsSymlink(symlinkname)

	// Check the result
	msg = fmt.Sprintf("The Symlink exist: [%v] Symlinkpath: [%v]", is_ok, symlinkname)
	showMsg(msg, is_ok, t)
}

// Function to test the IsFile
func TestIsFile(t *testing.T) {
	// Getting the local temp dirs
	tempDir := os.TempDir()

	// Create a new temp directory
	dir, err := ioutil.TempDir(tempDir, "isfile")
	checkError(err)

	// Remove the temp dir after the function is done
	defer os.RemoveAll(dir)

	// Create a temp file inside the temp directory
	file, err := ioutil.TempFile(dir, "")

	// Check if the object is a symlink
	is_ok := IsFile(file.Name())
	checkError(err)

	// Check the result
	msg = fmt.Sprintf("The File exist: [%v] Filepath: [%v]", is_ok, file.Name())
	showMsg(msg, is_ok, t)
}

// Function to test the IsDir
func TestIsDir(t *testing.T) {
	// Getting the local temp dirs
	tempDir := os.TempDir()

	// Create a new temp directory
	dir, err := ioutil.TempDir(tempDir, "isdir")
	checkError(err)

	// Remove the temp dir after the function is done
	defer os.RemoveAll(dir)

	// Check if the object is a directory
	is_ok = IsDir(dir)

	// Check if the directory exist
	msg = fmt.Sprintf("The Directory exist: [%v] Dirpath: [%v]", is_ok, dir)
	showMsg(msg, is_ok, t)
}

// Function to test the IsDir
func TestObjUsable(t *testing.T) {
	// Synlink Name
	var symlinkname string

	// Getting the local temp dirs
	tempDir := os.TempDir()

	// Create a new temp directory
	dir, err := ioutil.TempDir(tempDir, "isfile")
	checkError(err)

	// Remove the temp dir after the function is done
	defer os.RemoveAll(dir)

	// Create a temp file inside the temp directory
	file, err := ioutil.TempFile(dir, "")
	checkError(err)

	// Check if the file is usable or not
	is_ok = ObjUsable(file.Name())

	// Check the result
	msg = fmt.Sprintf("The File is usable: [%v] Filepath: [%v]", is_ok, file.Name())
	showMsg(msg, is_ok, t)

	// Check if the directory is usable or not
	is_ok = ObjUsable(dir)

	// Check the result
	msg = fmt.Sprintf("The Directory is usable: [%v] Dirpath: [%v]", is_ok, dir)
	showMsg(msg, is_ok, t)

	// Creating the symlink
	symlinkname = fmt.Sprintf("%v-%v", file.Name(), "link")
	os.Symlink(file.Name(), symlinkname)

	// Check if the symlink is usable or not
	is_ok = ObjUsable(symlinkname)

	// Check the result
	msg = fmt.Sprintf("The Symlink is usable: [%v] Dirpath: [%v]", is_ok, symlinkname)
	showMsg(msg, is_ok, t)
}

// Function to test the IsZeroOfUnderlyingTypes
func TestIsZeroOfUnderlyingType(t *testing.T) {
	// Needs to use the interface to hold the number 0
	isZero = 0

	// Check if the interface is igual zero (0)
	// It will not conver the string "0" to integer 0
	is_ok = IsZeroOfUnderlyingType(isZero)

	// Check the result
	msg = fmt.Sprintf("The Interface is equal Zero: [%v]", is_ok)
	showMsg(msg, is_ok, t)
}

// Function to test the ClearStr
func TestClearStr(t *testing.T) {
	stringWithSpaces := " String "
	stringWithTabs := "\tString"
	stringWithNewLine := "String\n"
	stringWithNewLineAndReturn := "String\r\n"

	// Check if the string has any space
	is_ok = strings.Contains(stringWithSpaces, " ")

	// Check the result
	msg = fmt.Sprintf("The String has Spaces: [%v]", is_ok)
	showMsg(msg, is_ok, t)

	// Check if the string has any space
	is_ok = !strings.Contains(ClearStr(stringWithSpaces), " ")

	// Check the result
	msg = fmt.Sprintf("The String has no Spaces: [%v]", is_ok)
	showMsg(msg, is_ok, t)

	// Check if the string has any tabs
	is_ok = strings.Contains(stringWithTabs, "\t")

	// Check the result
	msg = fmt.Sprintf("The String has Tabs: [%v]", is_ok)
	showMsg(msg, is_ok, t)

	// Check if the string has any tabs
	is_ok = !strings.Contains(ClearStr(stringWithTabs), "\t")

	// Check the result
	msg = fmt.Sprintf("The String has no Tabs: [%v]", is_ok)
	showMsg(msg, is_ok, t)

	// Check if the string has any new line
	is_ok = strings.Contains(stringWithNewLine, "\n")

	// Check the result
	msg = fmt.Sprintf("The String has NewLine: [%v]", is_ok)
	showMsg(msg, is_ok, t)

	// Check if the string has any new line
	is_ok = !strings.Contains(ClearStr(stringWithNewLine), "\n")

	// Check the result
	msg = fmt.Sprintf("The String has no NewLine: [%v]", is_ok)
	showMsg(msg, is_ok, t)

	// Check if the string has any new line and return
	is_ok = strings.Contains(stringWithNewLineAndReturn, "\r\n")

	// Check the result
	msg = fmt.Sprintf("The String has NewLine and Return: [%v]", is_ok)
	showMsg(msg, is_ok, t)

	// Check if the string has any new line and return
	is_ok = !strings.Contains(ClearStr(stringWithNewLineAndReturn), "\r\n")

	// Check the result
	msg = fmt.Sprintf("The String has no NewLine or Return: [%v]", is_ok)
	showMsg(msg, is_ok, t)

}

// Function to show the validation message
func showMsg(msg string, check bool, t *testing.T) {
	if !check {
		t.Errorf(msg)
	} else {
		t.Logf(msg)
	}
}

// Function to check if there is errorss
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
