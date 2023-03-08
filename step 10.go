package main

import (
	"fmt"
	"os"
)

func main() {
	// Obtiene los argumentos de entrada del programa
	args := os.Args

	// Verifica si hay tres argumentos de entrada
	if len(args) != 4 {
		fmt.Println("Debe ingresar tres argumentos de entrada")
		return
	}

	// Obtiene los argumentos de entrada individuales
	arg1 := args[1]
	arg2 := args[2]
	arg3 := args[3]

	// Imprime los argumentos de entrada
	fmt.Println("El primer argumento es:", arg1)
	fmt.Println("El segundo argumento es:", arg2)
	fmt.Println("El tercer argumento es:", arg3)
}
