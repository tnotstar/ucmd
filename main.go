package main

import (
	"fmt"
)

func main() {
	ReadEvalPrintLoop()
}

func ReadEvalPrintLoop() {
	for {
		fmt.Print(">>> ")

		var input string
		fmt.Scanln(&input)
		if input == "exit" {
			break
		}

		fmt.Printf("Hello, %v!\n", input)
	}
}
