package main

import "fmt"

func main() {
	// Operator Mathematics
	var (
		a = 10
		b = 10
		c = 5
		d = 7
		e = 12
		f = 4
		g = a + b - c * d / e % f
	)
	fmt.Println("The result of g is:", g)

	// Assignment Operator
	// var (
	// 	i = 10
	// 	j = 10
	// 	k = 5
	// 	l = 7
	// 	m = 3
	// )
	// 	i += 10 // i = i + 10
	// 	j -= 5 // j = j - 5
	// 	k *= 2 // k = k * 2
	// 	l /= 2 // l = l / 2
	// 	m %= 2 // m = m % 2

	// fmt.Println(i, j, k, l, m)

	// Unary Operator
	// var (
	// 	i = 10
	// 	j = 10
	// )
	// i++ // i = i + 1
	// j-- // j = j - 1
	// fmt.Println(i, j)
}
