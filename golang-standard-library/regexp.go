package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("faris"))
	fmt.Println(regex.MatchString("fariz"))
	fmt.Println(regex.MatchString("faRis"))

	fmt.Println(regex.FindAllString("faris fariz rifas masud f4ris sudmas faRis", 10))
}