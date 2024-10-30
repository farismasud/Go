package main

import (
	"fmt"
	"flag"
)
func main()  { 
	host := flag.String("host", "localhost", "Put your host")
	username := flag.String("username", "root", "Put your username")
	password := flag.String("password", "root", "Put your password")
	port := flag.Int("port", 8080, "Put your port")
	flag.Parse()

	fmt.Println(*host, *username, *password, *port)
}