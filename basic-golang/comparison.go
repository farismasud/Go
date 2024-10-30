package main

import "fmt"

func main() {
	var name1 = "Faris"
	var name2 = "Faris"
	var name3 = "Mas'ud"

	fmt.Println(name1 == name2) // true
	fmt.Println(name1 != name3) // true
	fmt.Println(name1 != name2) // false
}

