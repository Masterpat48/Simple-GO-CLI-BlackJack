package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var hand int
	dealerhand := rand.Intn(9) + 1
	fmt.Println("-----BLACK JACK-----")
	for i := 0; i >= 0; i++ {
		hand += rand.Intn(9) + 1
		fmt.Println("New hand: ", hand)

		if hand == 21 {
			fmt.Println("BLACKJACK - You won!")
			fmt.Scanln()
			os.Exit(1)
		} else if hand > 21 {
			fmt.Println("You lost")
			fmt.Scanln()
			os.Exit(1)
		}

		fmt.Println("You have ", hand, ", do you want to HIT or STAND")

		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case "HIT":
			continue
		case "STAND":
			for j := 0; j >= 0; j++ {

				fmt.Println("Dealer hand: ", dealerhand)
				if dealerhand < 17 {
					fmt.Println("Dealer hand is lower than 17 so he's gonna HIT")
					dealerhand = dealerhand + (rand.Intn(9) + 1)
					continue
				} else if dealerhand > 21 {
					fmt.Println("DEALER BUSTED - You won!")
					fmt.Scanln()
					os.Exit(1)
				} else {
					break
				}
			}

			if hand < 21 && hand > dealerhand {
				fmt.Println("You won!")
				fmt.Scanln()
				os.Exit(1)
			} else if hand <= 21 && hand == dealerhand {
				fmt.Println("Push!")
				fmt.Scanln()
				os.Exit(1)
			} else if hand == 21 {
				fmt.Println("BLACKJACK - You won!")
				fmt.Scanln()
				os.Exit(1)
			} else {
				fmt.Println("You lost")
				fmt.Scanln()
				os.Exit(1)
			}
		}

	}
}
