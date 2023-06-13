package main

import "fmt"

func main() {
	var a string
	fmt.Print("Enter your Name:")
	fmt.Scanln(&a)
	fmt.Print("my name is", a)

}
