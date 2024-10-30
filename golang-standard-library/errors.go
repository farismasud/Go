package main

import "fmt"

import "errors"

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "faris" {
		return NotFoundError
	}
	//success
	return nil
}

func main() {
	err := GetById("mas'ud")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found")
		} else {
			fmt.Println("unknown error")
		}
	}
}
