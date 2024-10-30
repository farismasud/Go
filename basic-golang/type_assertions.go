package main

import "fmt"

func random() any {
	// return "OK"
	// return 100
	return true
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type){
	case string :
		fmt.Println("String", value)
	case int :
		fmt.Println("Int", value)
	default	:
		fmt.Println("Uknown", value)
	}
}
