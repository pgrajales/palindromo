package main

import (
	"ProjectUQH/src"
	"fmt"
)

func main() {
	//18446744073709551615
	var eleccion int64
	fmt.Scanln(&eleccion)
	output := src.IsPalindrome(eleccion)
	fmt.Printf("El número %v es un palindromo?: %v", eleccion, output)

}
