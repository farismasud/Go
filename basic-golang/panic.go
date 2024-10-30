package main

import "fmt"

func endApp(){
	fmt.Println("End app")
	message := recover()
	fmt.Println("Terjadi Panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}
}

func main(){
	runApp(true)
	fmt.Println("Faris Mas'ud")
}