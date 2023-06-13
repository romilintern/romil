package main

import (
	"fmt"
	"thisistest/calculator"
)

func main() {
	var a float64
	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)
	var b float64
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)
	var c string
	fmt.Println("Enter the choice for what you want to perform")
	fmt.Println("+ : Addition")
	fmt.Println("- : Subraction")
	fmt.Println("* : Multiplication")
	fmt.Println("/ : Division")
	fmt.Scanln(&c)
	var result float64
	result = calculator.Calci(c, a, b)
	fmt.Print(result)
}
