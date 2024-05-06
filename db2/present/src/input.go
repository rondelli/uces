package main

import (
	"fmt"
)

func main() {
	var nombre string

	fmt.Printf("Ingresá tu nombre: ")
	fmt.Scanf("%s", &nombre) // HL

	fmt.Printf("Hola, %s!\n", nombre)
}
