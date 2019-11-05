package main

/*
This program take a number entered from user
and compare it to random system number
*/
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numberOfAttemps := 3
	userNumber := 0
	winner := false
	// rand.Seed you'll receive different results at each execution
	rand.Seed(time.Now().UnixNano())
	for numberOfAttemps != 0 {
		randNumber := rand.Intn(11)
		println("Please guess some  number between  0 to 10")
		fmt.Printf("You have now %d attemps\n", numberOfAttemps)
		fmt.Scanf("%d", &userNumber)
		if userNumber == randNumber {
			println("Congratulations you won!!!")
			winner = true
			break
		}
		fmt.Printf("): The system number was %d\n", randNumber)
		numberOfAttemps--
	}
	if !winner {
		println("oh no,  you lost")
	}
}
