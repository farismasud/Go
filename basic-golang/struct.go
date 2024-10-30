package main

import "fmt"

type Customer struct{
	Name 		string
	Address		string
	Age			int
}

func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main(){
	var faris Customer
	faris.Name = "Faris Mas'ud"
	faris.Address = "Indonesia"
	faris.Age = 23

	fmt.Println(faris)
	fmt.Println(faris.Name)
	fmt.Println(faris.Address)
	fmt.Println(faris.Age)


	//Struct Literal
	masud := Customer{
		Name: 		"Mas'ud",
		Address: 	"Indonesia",
		Age:		30,
	}
	fmt.Println(masud)

	// OR
	reva := Customer{"Masud", "Indonesia", 30}
	fmt.Println(reva)

	reva.sayHello("Masud")
}