package main

import (
	"time"
	"fmt"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now)
	fmt.Println(now.Local())

	var utc time.Time = time.Date(2024, time.May, 17, 5, 1, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	
	value := "2024-05-17 05:01:00"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else{
		fmt.Println(valueTime)
	}
	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
	fmt.Println(valueTime.Minute())
	fmt.Println(valueTime.Second())

}