package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//! Pointer By Value
	// address1 := Address{"Jakarta,", "DKI Jakarta,", "Indonesia"}
	// address2 := address1 //! By Value

	// address2.City = "Depok,"
	// address2.Province = "Jawa Barat"

	// fmt.Println(address1)
	// fmt.Println(address2)

	//! Pointer by Reference
	var address1 = Address{"Jakarta,", "DKI Jakarta,", "Indonesia"}
	var address2 *Address = &address1 //! By Reference

	address2.City = "Depok,"
	address2.Province = "Jawa Barat"

	fmt.Println(address1)
	fmt.Println(address2)
}
