package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10{
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	//}

	// For with Statement
	for counter := 1; counter <= 10; counter++{
		fmt.Println("Perulangan ke", counter)
	}
	fmt.Println("Finish")

	//For Range Manual
	names := []string{"Faris", "Mas'ud", "Faris Mas'ud"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	//For Range 
	for index, name := range names{
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names{
		fmt.Println(name)
	}
}