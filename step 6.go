package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
				fmt.Print("Se ha agregado la imagen: ", nombres[i], "\n")
				i++
			}
		}
	}
	//fmt.Println(nombres)
	fmt.Println("Cantidad de imagenes encontradas: ", len(nombres))
}

//https://steemit.com/cervantes/@orlmicron/array-y-slice-en-go-golang
