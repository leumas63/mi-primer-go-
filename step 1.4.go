package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	archivos, err := ioutil.ReadDir("D:\\Users\\Cristian\\Documents\\2023-1\\Computacion Nube\\Pruebas")
	if err != nil {
		log.Fatal(err)
	}
	for _, archivo := range archivos {
		fmt.Println("Nombre:", archivo.Name())
		fmt.Println("\n")
	}
}
