package checks

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// Function to test the IsSymlink
func TestIsSymlink(t *testing.T) {

	// Message will be hold here
	var msg string

	// Synlink Name
	var symlinkname string

	// Getting the local temp dirs
	tempDir := os.TempDir()

	// Create a new temp directory
	dir, err := ioutil.TempDir(tempDir, "isfile")
	if err != nil {
		log.Fatal(err)
	}

	// Remove the temp dir after the function is done
	defer os.RemoveAll(dir)

	// Create a temp file inside the temp directory
	file, err := ioutil.TempFile(dir, "")
	if err != nil {
		log.Fatal(err)
	}

	// Creating the symlink pointing to the tempfile
	symlinkname = fmt.Sprintf("%v-%v", file.Name(), "link")
	os.Symlink(file.Name(), symlinkname)

	// Check if the object is a symlink
	ok := IsSymlink(symlinkname)

	// Check the result
	msg = fmt.Sprintf("The Synlink exist: [%v] synlinkpath: [%v]", ok, symlinkname)
	if !ok {
		t.Errorf(msg)
	} else {
		t.Logf(msg)
	}
}

// Function to test the IsFile
func TestIsFile(t *testing.T) {

	// Message will be hold here
	var msg string

	// Getting the local temp dirs
	tempDir := os.TempDir()

	// Create a new temp directory
	dir, err := ioutil.TempDir(tempDir, "isfile")
	if err != nil {
		log.Fatal(err)
	}

	// Remove the temp dir after the function is done
	defer os.RemoveAll(dir)

	// Create a temp file inside the temp directory
	file, err := ioutil.TempFile(dir, "")

	// Check if the object is a symlink
	ok := IsFile(file.Name())
	if err != nil {
		log.Fatal(err)
	}

	// Check the result
	msg = fmt.Sprintf("The file exist: [%v] filepath: [%v]", ok, file.Name())
	if !ok {
		t.Errorf(msg)
	} else {
		t.Logf(msg)
	}
}

// Function to test the IsDir
func TestIsDir(t *testing.T) {

	// Message will be hold here
	var msg string

	// Getting the local temp dirs
	tempDir := os.TempDir()

	// Create a new temp directory
	dir, err := ioutil.TempDir(tempDir, "isdir")
	if err != nil {
		log.Fatal(err)
	}

	// Remove the temp dir after the function is done
	defer os.RemoveAll(dir)

	// Check if the object is a directory
	ok := IsDir(dir)

	// Check if the directory exist
	msg = fmt.Sprintf("The directory exist: [%v] dirpath: [%v]", ok, dir)
	if !ok {
		t.Errorf(msg)
	} else {
		t.Logf(msg)
	}
}

// Function to test the IsDir
func TestObjUsable(t *testing.T) {

	// Message will be hold here
	var msg string

	// Check if it usable
	var usable bool

	// Synlink Name
	var symlinkname string

	// Getting the local temp dirs
	tempDir := os.TempDir()

	// Create a new temp directory
	dir, err := ioutil.TempDir(tempDir, "isfile")
	if err != nil {
		log.Fatal(err)
	}

	// Remove the temp dir after the function is done
	defer os.RemoveAll(dir)

	// Create a temp file inside the temp directory
	file, err := ioutil.TempFile(dir, "")
	if err != nil {
		log.Fatal(err)
	}

	// Check if the file is usable or not
	usable = ObjUsable(file.Name())

	// Check the result
	msg = fmt.Sprintf("The file is usable: [%v] filepath: [%v]", usable, file.Name())
	if !usable {
		t.Errorf(msg)
	} else {
		t.Logf(msg)
	}

	// Check if the directory is usable or not
	usable = ObjUsable(dir)

	// Check the result
	msg = fmt.Sprintf("The directory is usable: [%v] dirpath: [%v]", usable, dir)
	if !usable {
		t.Errorf(msg)
	} else {
		t.Logf(msg)
	}

	// Creating the symlink
	symlinkname = fmt.Sprintf("%v-%v", file.Name(), "link")
	os.Symlink(file.Name(), symlinkname)

	// Check if the symlink is usable or not
	usable = ObjUsable(symlinkname)

	// Check the result
	msg = fmt.Sprintf("The symlink is usable: [%v] dirpath: [%v]", usable, symlinkname)
	if !usable {
		t.Errorf(msg)
	} else {
		t.Logf(msg)
	}

}
