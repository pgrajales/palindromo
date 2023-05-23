package main

import (
	"ProjectUQH/src"
	"fmt"
)

func main() {
	var eleccion int
	fmt.Scanln(&eleccion)
	output := src.IsPalindrome(eleccion)
	fmt.Printf("El número %v es un palindromo?: %v", eleccion, output)

}
