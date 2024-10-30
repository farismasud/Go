package main

import "fmt"

func getGoodBye(name string) string{
	return "Good Bye " + name
}

func main()  {
	contoh := getGoodBye //function as values(dont use ())

	fmt.Println(contoh("Faris"))
}