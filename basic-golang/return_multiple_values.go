package main

import "fmt"

func getFullName()(string, string)  {
	return "Faris", "Mas'ud"
}

func main()  {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}

//! Ignore Return
// func main()  {
// 	firstName, _ := getFullName()
// 	fmt.Println(firstName)
// }