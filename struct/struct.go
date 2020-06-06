package main

import "fmt"

type user struct {
	name string
	age  int8
	city string
}

func main() {
	u := []user{}

	// for i := range u{
	// 	fmt.Println("Enter Your Name And Age");
	// 	n, _ := fmt.Scanf("%s\n%d\n", &u[i].name, &u[i].age)
	// 	fmt.Println(u[i],n)
	// }
		u=append(u,user{"Hawk",12,"LA"})
		fmt.Printf("%T",u[0].name);
}
