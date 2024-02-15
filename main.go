package main

import (
	"fmt"
	"strings"
)

func getUserInput() string {
	fmt.Println("Who is the best person in this room?")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return input
}

func printBestPerson(bestPerson string) {
	bestPerson = strings.ToLower(bestPerson)
	if bestPerson == "joe" {
		fmt.Println("I mean I agree but sucking up to the creator of this program won't get you anywhere")
	} else if bestPerson == "molly" {
		fmt.Println("I take back what I said about Joe, Molly is the best person in this room, still doesn't mean you get a hint")
	} else {
		fmt.Println(bestPerson, "is the best person in this room")
	}
}

func main() {
	bestPerson := getUserInput()
	printBestPerson(bestPerson)
}
