package main

import "fmt"

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
	fmt.Println(bestPerson, "is the best person in this room")
}

func main() {
	bestPerson := getUserInput()
	printBestPerson(bestPerson)
}
