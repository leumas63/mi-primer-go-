package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"strings"
)

func main() {
	var nombres []string
	var i int
	archivos, err := ioutil.ReadDir("D:\\Users\\Cristian\\Documents\\2023-1\\Computacion Nube\\Pruebas")
	if err != nil {
		log.Fatal(err)
	}
	for _, archivo := range archivos {
		if !archivo.IsDir() {
			ext := strings.ToLower(filepath.Ext(archivo.Name()))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {

				nombres = append(nombres, archivo.Name())
				i++
			}
		}
	}
	fmt.Println("Cantidad de imagenes encontradas: ", len(nombres))

	//[] Aleatorio
	length := len(nombres)
	aleatorio := mostarAleatorio(length)
	fmt.Println("La imagen aleatoria es: ", nombres[aleatorio])
}

func mostarAleatorio(k int) int {
	var numero int
	numero = rand.Intn(k)

	return numero
}
