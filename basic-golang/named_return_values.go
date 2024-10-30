package main

import "fmt"

func getComplateName()(firstName, midleName, lastName string)  {
	firstName = "faris"
	midleName = "mas'ud"
	lastName = "Faris Mas'ud"

	return firstName, midleName, lastName
}

func main()  {
	a, b, c := getComplateName()
	fmt.Println(a, b, c)
}