package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Print("Enter your Name:")
	fmt.Scanln(&a)
	capital(a)
}
func capital(a string) {
	a = strings.ToUpper(a)
	print("Name in Capital is: ", a)

}
