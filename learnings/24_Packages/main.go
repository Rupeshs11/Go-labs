package main // main package is the entry point

import (
	"fmt"

	"github.com/rupeshs11/Go-labs/learnings/auth"
	"github.com/rupeshs11/Go-labs/learnings/user"
	"github.com/fatih/color"
) // packages are use for best practices of software engineering

// using package we can speed up the compilation time
func main() {
	auth.LoginWithCredentials("Knox", "Joyboy123")

	session := auth.GetSession()
	fmt.Println("session:", session)

	user := user.User{
		Email: "user@email.com",
		Name:  "Monkey D. Luffy",
	}
	//fmt.Println(user.Email,user.Name)
	color.Magenta(user.Email)
	color.Green(user.Name)
}

// if we want to make somthing private make the first letter small else for public it must be capital

//go command for go.mod : go mod init github.com/rupeshs11/Go-labs/learnings
// to install the thirdparty package we use  go get github.com/username/packagename
// to fix the important things into package suse go mod tidy
