package main

import "fmt"

func main() {
	const firstName string = "Faris"
	const lastName = "Mas'ud"

	//multiple const
	// const (
	// 		firstName string = "Faris"
	// 		lastName string = "Mas'ud"
	//)

	//error cannot change if it is const
	//firstName = "Mas'ud"
	//lastName = "Faris"

	fmt.Println(firstName, lastName)
}

