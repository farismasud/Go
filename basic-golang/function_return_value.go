package main

import "fmt"

func getHelp(name string) string {
	return "Help " + name
}

func main() {
	result := getHelp("Faris")
	fmt.Println(result)
}