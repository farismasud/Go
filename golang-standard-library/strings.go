package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Faris Mas'ud", "Faris"))
	fmt.Println(strings.Split("Faris Mas'ud", " "))
	fmt.Println(strings.ToLower("Faris Mas'ud"))
	fmt.Println(strings.ToUpper("faris mas'ud"))
	fmt.Println(strings.Trim("   Faris Mas'ud   ", " "))
	fmt.Println(strings.ReplaceAll("Faris Mas'ud", "Mas'ud", "Faris"))
}