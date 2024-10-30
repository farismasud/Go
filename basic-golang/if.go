package main

import "fmt"

func main(){
	name := "Faris"
	
	if name == "Faris" {
		fmt.Println("Hello Faris")
	} else if name == "Mas'ud" {
		fmt.Println("Hello Mas'ud")
	} else{
		fmt.Println("Hello, What it's your name?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Name is too long")
	} else {
		fmt.Println("Name are correct")
	}
}