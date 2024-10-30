package main

import (
	"basic-golang/database"
	_ "basic-golang/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
