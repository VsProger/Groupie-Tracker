package main

import (
	"fmt"
	"groupie-tracker/cmd"
)

func main() {
	fmt.Println("Server available on http://localhost:8080/")
	cmd.StartServer()
}
