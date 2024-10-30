package main

import "fmt"

func main()  {
	 name := "Faris"

	switch name {
	case "Faris":
		fmt.Println("Halo" ,name)
	case "Mas'ud":
		fmt.Println("Hai", name)
	default:
		fmt.Println("Hello, What it's your name?")
	}

	//* Switch Short 
	switch length := len(name); length > 5{
	case true:
		fmt.Println("Name is too long")
	case false:
		fmt.Println("Name are correct")
	}


	//? Switch without Condition
	length := len(name)
	switch {
	case length > 6:
		fmt.Println("Name is too long")
	case length < 6:
		fmt.Println("Name are correct")
	default:
		fmt.Println("Hello, What it's your name?")
	}
}