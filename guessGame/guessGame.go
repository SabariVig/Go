package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //Initilizes A Random Number Every Time Else Go Seeds The Rand Func At Complie Time
	C := rand.Intn(10)
	fmt.Print(C)

	for i := 1; ; i++ {
		inp := 0
		fmt.Print("Enter A Number: ")
		fmt.Scan(&inp)

		if inp < C {
			fmt.Println("Try A Greater Number")
		} else if inp > C {
			println("Try A Smaller Number")
		} else {
			fmt.Println("You Gussed In", i, "Tries")
			break

		}

	}

}
