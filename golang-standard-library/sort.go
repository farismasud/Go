package main

import (
	"sort"
	"fmt"
)

type User struct {
	name string
	age int
}

// Contract
type UserSlice []User

func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].age < u[j].age
}

func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User{
		{"faris", 22},
		{"mas'ud", 23},
		{"reva", 21},
		{"amelia", 24},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}