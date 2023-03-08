package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// Obtiene el nombre del host del equipo
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	// Imprime el nombre del host del equipo
	fmt.Println("Nombre de host: ", hostname)
}
