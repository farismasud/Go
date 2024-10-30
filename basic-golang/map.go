package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Faris",
		"address": "Depok",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	zoo := make(map[string]string)
	zoo["Name"] = "Leonidas"
	zoo["Spesies"] = "Leo"
	zoo["ups"] = "Wrong"

	fmt.Println(zoo)

	delete(zoo, "ups")
	fmt.Println(zoo)
}
