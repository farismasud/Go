package helper

var version = "1.0.0"
var Application = "Basic Golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh(){
	sayGoodBye("Faris")
	fmt.Println(version)
}