package main

import "fmt"

func getUserInput() {
	fmt.Println("Who is the best person in this room?")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("You think", input, "is the best person in this room.")
}

func main() {
	getUserInput()
}
