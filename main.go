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
	switch bestPerson {
	case "joe":
		fmt.Println("I mean I agree but sucking up to the creator of this program won't get you anywhere")
	case "molly":
		fmt.Println("I take back what I said about Joe, Molly is the best person in this room, still doesn't mean you get a hint")
	case "joey":
		fmt.Println("Joey may be the current chairperson but not even he knows this flag!")
	case "izzy":
		fmt.Println("I didn't think anyone would think Izzy was the best person, unless you have a crush on her of course, but I guess you can have a flag she would like dmu_hackers{IzzyIsTheBest}. It won't work but you can have it")
	case "orlin":
		fmt.Println("I didn't know Maisie was in this room. Live orlin update anyone?")
	case "maisie":
		fmt.Println("Didn't think Maisie was here, but I guess you're just trying to be Orlin")
	case "dan":
		fmt.Println("Either you're looking for miso or DJ, either way, you're not getting a flag")
	case "miso":
		fmt.Println("Thank you for specifying the correct Dan. But no flag for you")
	case "dj":
		fmt.Println("Thank you for specifying the correct Dan. But no points for being named Emma")
	case "emma":
		fmt.Println("Didn't think Emma came to hacker's sessions")
	default:
		fmt.Println(bestPerson, "is the best person in this room")
	}
}

func main() {
	bestPerson := getUserInput()
	printBestPerson(bestPerson)
}
