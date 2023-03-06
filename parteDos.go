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
		fmt.Println("Tamaño:", archivo.Size())
		fmt.Println("Modo:", archivo.Mode())
		fmt.Println("Ultima modificación:", archivo.ModTime())
		fmt.Println("Es directorio?:", archivo.IsDir())
		fmt.Println("\n")
	}
}

//https://apuntes.de/golang/manejo-de-archivos-listar-un-directorio/#gsc.tab=0
/*package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "strings"
)

func main() {
    dir := "../imagenes" // Reemplazar con la ruta al directorio deseado

    files, err := ioutil.ReadDir(dir)
    if err != nil {
        panic(err)
    }

    for _, file := range files {
        if !file.IsDir() {
            ext := strings.ToLower(filepath.Ext(file.Name()))
            if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
                fmt.Println(file.Name())
            }
        }
    }
}*/
