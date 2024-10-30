package main

import "fmt"

// basic Functions
// func sayHelp(){
// 	fmt.Println("Help")
// }

// func main()  {
// 	sayHelp()
// }

// Functions Params
func sayHelpTo(firstName string, lastName string) {
	fmt.Println("Help", firstName, lastName)
}

func main(){
	sayHelpTo("Faris", "Mas'ud")
}