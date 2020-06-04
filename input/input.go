package main

import "fmt"

func main() {
	fmt.Println("Hello")
	c := 0 // Initialize variable to accept integer input
	fmt.Println("Enter Your Birth Year")
	fmt.Scan(&c)
	if c >= 1 {
		fmt.Println("Your Age Will Be", 2020-c, "At The End Of 2020")
	}

}
