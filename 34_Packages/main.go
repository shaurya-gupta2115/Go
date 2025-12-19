package main

import (
	"fmt"

	"github.com/shaurya-gupta2115/Go/34_Packages/auth"
	"github.com/shaurya-gupta2115/Go/34_Packages/user"
)

func main() {
	auth.LoginWithCredentials("shauryagupta", "1234567890")

	a := auth.GetSession()
	fmt.Println(a)

	//importing from user package
	newUser := user.User{
		Name:  "Shaurya Gupta",
		Email: "vishnu8858369983@gmail.com",
	}
	fmt.Println(newUser)
	fmt.Println(newUser.Email)
	fmt.Println(newUser.Name)

	// color.Red("Teri maa ki chut") // this is how we use the packages

}
