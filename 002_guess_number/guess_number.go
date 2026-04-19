package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var win int = 0
	var lose int = 0
	var try int = 0
	
	var min int = 1
	var max int = 5


	fmt.Println("Welcome to Guessing number game!")
	fmt.Printf("Guess number from %d to %d:\n", min, max)
	fmt.Println("Push Ctrl + C to stop game")

	for {
		try += 1
		var userAnswer int
		var randomNumber int = rand.Intn(max - min + 1) + min

		fmt.Printf("==== Try №%d ====\n", try)
		fmt.Print("Enter number: ")
		
		_, err := fmt.Scanln(&userAnswer)

		if err != nil{
			fmt.Println("Please, type a valid number")
			fmt.Scanln()			
			continue
		}

		if userAnswer == randomNumber{
			win +=1			
			fmt.Printf("Great! %d was guessed\n", randomNumber)			
		} else if(userAnswer < min || userAnswer > max) {
			fmt.Printf("Please, type number between %d and %d\n", min, max)
			
		} else {
			lose += 1 			
			fmt.Printf("You lose! %d was guessed\n", randomNumber)			
		}

		fmt.Printf("Stats: Try %d | Wins %d | Loses: %d\n", try, win, lose)
		fmt.Print("\n")
				
	}

}