package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Hasil", result)
	}

	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Hasil", resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println("Binary", binary)

	stringInt := strconv.Itoa(1000)
	fmt.Println(stringInt)
}