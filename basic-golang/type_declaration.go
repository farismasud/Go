package main

import "fmt"

func main() {
	type NoKTP string

	var ktpFaris NoKTP = "1234567890"

	var contoh NoKTP = "0987654321"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpFaris)
	fmt.Println(contohKtp)
}

