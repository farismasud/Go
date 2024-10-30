package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main()  {
	faris := Man{"Faris"}
	faris.Married()

	fmt.Println(faris.Name)
}
