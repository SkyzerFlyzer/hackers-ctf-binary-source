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
	if len(bestPerson) > 20 {
		message := " is the best person in this room"
		post20 := bestPerson[20:]
		for i := 0; i < len(post20) && i < len(message); i++ {
			bestPerson = bestPerson[:20] + string(message[i])
		}
		fmt.Println(bestPerson)
		return
	}
	switch bestPerson {
	case "joe":
		fmt.Println("I mean I agree but sucking up to the creator of this program won't get you anywhere")
	case "molly":
		fmt.Println("I take back what I said about Joe, Molly is the best person in this room, still doesn't mean you get a hint")
	case "joey":
		fmt.Println("Joey may be the current chairperson but not even he knows this flag!")
	case "izzy":
		fmt.Println("I didn't think anyone would think Izzy was the best person, unless you have a crush on her of course, but I guess you can have a flag she would like dmu_hackers{IzzyIsTheBest}. It won't work as it's a complete lie but you can have it")
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
	case "matthew":
		fmt.Println("Matthew or is it Matt? Must be one of those new committee members.")
	case "matt":
		fmt.Println("Matthew or is it Matt? Must be one of those new committee members.")
	case "rory":
		fmt.Println("Roary the racing car, Roary the number 1 star!")
	case "jack":
		fmt.Println("Will Jack drink alcohol or just play chess? Find out next time on Dragon Ball Z")
	case "jordon":
		fmt.Println("I don't know if Jordon will be here but I'm sure he'd try a buffer overflow. Yes it is spelt with an O contrary to autocorrect's beliefs")
	case "chloe":
		fmt.Println("I think you meant to type Jack as we all know Chloe is a myth")
	case "george":
		fmt.Println("This man is good for all your tech support needs, the SU hired him for a reason")
	case "dilan":
		fmt.Println("Yes hes on placement but don't stay at his house overnight, you may get a full report in the morning")
	case "georgie":
		fmt.Println("She needs a bit of a break, being a nurse and all.")
	case "rav":
		fmt.Println("Someone get this man a high paying job that doesn't require security checks")
	default:
		fmt.Println(bestPerson, "is the best person in this room")
	}
}

func displayRepoInfo() {
	fmt.Println("The source code for the program you are currently running can be found at https://github.com/SkyzerFlyzer/hackers-ctf-binary-source")
}

func main() {
	displayRepoInfo()
	bestPerson := getUserInput()
	printBestPerson(bestPerson)
}
