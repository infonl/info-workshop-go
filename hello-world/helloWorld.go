package helloWorld

import "fmt"

const Text = "Hello World!"

func GetHelloWorld() string {
	return Text
}
func GetHelloName(name string) string {
	return "Hello" + " " + name
}
func PrintHelloWorld() {
	fmt.Print("Hello World!")
}
