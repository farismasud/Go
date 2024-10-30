package main

import "fmt"

func main() {
	// array string
	// var names [4]string
	// names[0] = "Faris"
	// names[1] = "Mas'ud"
	// names[2] = "Reva"
	// names[3] = "Amelia"

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])
	// fmt.Println(names[3])

	// direct array (horizontal)
	// var values = [4]int{10, 20, 30, 40}
	// fmt.Println(values)

	// direct array (vertical)
	// var values = [4]int{
	// 	10,
	// 	20,
	// 	30,
	// 	40,
	// }
	// fmt.Println(values)

	//function array
	var values = [...]int{
		10,
		20,
		30,
		40,
	}
	fmt.Println(values)
	fmt.Println(len(values))
}