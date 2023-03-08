package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	archivos, err := ioutil.ReadDir("D:\\Users\\Cristian\\Documents\\2023-1\\Computacion Nube\\Pruebas")
	if err != nil {
		log.Fatal(err)
	}
	for _, archivo := range archivos {
		if !archivo.IsDir() {
			ext := strings.ToLower(filepath.Ext(archivo.Name()))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
				fmt.Println("Nombre:", archivo.Name())
				fmt.Println("Tamaño:", archivo.Size())
				fmt.Println("Modo:", archivo.Mode())
				fmt.Println("Ultima modificación:", archivo.ModTime())
				fmt.Println("Es directorio?:", archivo.IsDir())
				fmt.Println("\n")
			}
		}
	}
}
