package main

import (
	"encoding/base64"
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
	direccion := "D:\\Users\\Cristian\\Documents\\2023-1\\Computacion Nube\\Pruebas"
	archivos, err := ioutil.ReadDir(direccion)
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
	imagen_aleatoria := direccion + "\\" + nombres[aleatorio]

	//[] Codificación
	byte_imagen, err := ioutil.ReadFile(imagen_aleatoria)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(imagen_aleatoria)
	img_64 := base64.StdEncoding.EncodeToString(byte_imagen)
	fmt.Println(img_64)

}

func mostarAleatorio(k int) int {
	var numero int
	numero = rand.Intn(k)

	return numero
}
