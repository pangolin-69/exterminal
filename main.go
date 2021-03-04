package main

import (
	"exterminal/prompter"
	"fmt"
)

func main() {
	username, password := prompter.Credentials()
	fmt.Printf("Username: %s, Password: %s\n", username, password)
}
