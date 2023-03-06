package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	archivos, err := ioutil.ReadDir("C:\\Users\\pc\\Desktop\\8\\Computacion en la nube\\mi-primer-go\\imagenes")
	if err != nil {
		log.Fatal(err)
	}
	for _, archivo := range archivos {
		fmt.Println("Nombre:", archivo.Name())
		fmt.Println("\n")
	}
}
