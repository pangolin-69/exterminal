package main

import (
	"fmt"
	"github.com/pangolin-69/exterminal/src/prompter"
)

func main() {
	username, password := prompter.Credentials()
	fmt.Printf("Username: %s, Password: %s\n", username, password)
}
