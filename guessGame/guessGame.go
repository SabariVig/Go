package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	C := rand.Intn(10)
	fmt.Print(C)
	for i := 1; ; i++ {
		inp := 0
		fmt.Print("Enter A Number: ")
		fmt.Scan(&inp)
		if inp < C {
			fmt.Print("Try A Greater Number\n")
		}
		if inp > C {
			print("Try A Smaller Number\n")
		}
		if inp == C {
			fmt.Println("You Gussed In", i, "Number Of Times")
			break

		}

	}

}
