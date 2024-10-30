package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 = Address{"Jakarta,", "DKI Jakarta,", "Indonesia"}
	var address2 *Address = &address1 //! Pointer By Reference

	address2.City = "Bandung,"


	fmt.Println(address1)
	fmt.Println(address2)

	//address2 = &Address{"Depok,", "Jawa Barat,", "Indonesia"}
	*address2 = Address{"Depok,", "Jawa Barat,", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
