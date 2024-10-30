package main

import (
	"fmt"
	"basic-golang/helper"
)

func main() {
	result := helper.SayHello("Faris")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(helper.version)
	fmt.Println(helper.sayGoodBye("Faris"))
}