# Golang - Utils

[![Made with Go](https://img.shields.io/badge/go->%3D1.16-blue?style=flat-square&logo=go&logoColor=white)](https://golang.org)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/flownerd/utils)](https://pkg.go.dev/github.com/flownerd/utils)

[![Build](https://github.com/flownerd/utils/workflows/build/badge.svg)](https://github.com/flownerd/utils/actions)
[![CodeQL](https://github.com/flownerd/utils/workflows/CodeQL/badge.svg)](https://github.com/flownerd/utils/actions?query=workflow%3ACodeQL "Code quality workflow status")
[![GitHub tag](https://img.shields.io/github/v/tag/flownerd/utils.svg?sort=semver)](https://github.com/flownerd/utils/releases/?include_prereleases&sort=semver "View GitHub releases")
[![License](https://img.shields.io/badge/License-GPL%20v3-blue)](#license "Go to license section")

[![FlowNerd - utils](https://img.shields.io/static/v1?label=flownerd&message=utils&color=blue&logo=github)](https://github.com/flownerd/utils)
[![stars - utils](https://img.shields.io/github/stars/flownerd/utils?style=social)](https://github.com/flownerd/utils)
[![forks - utils](https://img.shields.io/github/forks/flownerd/utils?style=social)](https://github.com/flownerd/utils)

## Using the Input functions

Those functions were created to deal with the user input, by default we need to create the reader and after that deal with the newline returns and all the other issues that we can have, in that case I create some wrapper to deal with that.

- ReadString
  - Will return a string type
- ReadInt
  - Will return a int type
- ReadFloat
  - Will return a float64 type
- ReadBool
  - Will return a bool type

Create a new directory

```bash
mkdir sandbox && cd sandbox
```

Now we need to create a main.go file and the mod

```bash
touch main.go
go mod init sandbox
```

Now let's add the require module inside go.mod

```bash
module sandbox

go 1.17

require github.com/flownerd/utils v0.0.2 // Here I am using the version v0.0.2 if you want to use the lastest just remove the version
```

Now we can put the following content inside the main.go to test the utils module

```golang
package main

import (
 "fmt"
 inputs "github.com/flownerd/utils/functions/inputs"
)

// Structs
type User struct {
 UserName       string
 Age            int
 FavoriteNumber float64
 OwnsADog       bool
}

func main() {
 var user User
 user.UserName = inputs.ReadString("What is your name?\n")
 user.Age = inputs.ReadInt("How old are you?\n")
 user.FavoriteNumber = inputs.ReadFloat("What is your favorite number?\n")
 user.OwnsADog = inputs.ReadBool("Do you own a dog (y/n)?\n")
 fmt.Printf("Your name is %s, and you are %d years old. Your favourite number is %.2f. Owns a dog: %t.\n", user.UserName, user.Age, user.FavoriteNumber, user.OwnsADog)
}
```

Now we need to get the util modules and its dependencies

```bash
go get .
```

Now we can run the program

```bash
go run main.go
What is your name?
-> Douglas
How old are you?
-> 33
What is your favorite number?
-> 14.15
Do you own a dog (y/n)?
-> Your name is Douglas, and you are 33 years old. Your favourite number is 14.15. Owns a dog: true.
```
