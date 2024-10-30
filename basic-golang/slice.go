package main

import "fmt"

func main() {
	// slice := [...]string{"Faris", "Mas'ud", "Kamal", "John", "Doe", "Jacob"}
	// slice1 := slice[4:6]
	// fmt.Println(slice1)

	// slice2 := slice[:]
	// fmt.Println(slice2)

	// slice3 := slice[1:4]
	// fmt.Println(len(slice3))

	// slice4 := slice[2:]
	// fmt.Println(cap(slice4))

	// Append Slice
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	daySlice1 := days[5:] // Saturday and Sunday
	fmt.Println(daySlice1)

	daySlice1[0] = "New Saturday"
	daySlice1[1] = "New Sunday"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "New Holiday")
	daySlice2[0] = "Last Saturday"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// Make Slice
	// var newSlice []string = make([]string, 2, 5)
	// newSlice[0] = "Faris"
	// newSlice[1] = "Faris"
	// newSlice = append(newSlice, "Mas'ud") // Append
	// newSlice[2] = "Mas'ud"

	// fmt.Println(newSlice)
	// fmt.Println(len(newSlice))
	// fmt.Println(cap(newSlice))

	// Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Difference between Array and Slice
	// thisArray := [...]int{1, 2, 3, 4, 5} 
	// thisSlice := []int{1, 2, 3, 4, 5}

	// fmt.Println(thisArray)
	// fmt.Println(thisSlice)
}
