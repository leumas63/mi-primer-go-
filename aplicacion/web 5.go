package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Definir la estructura base del HTML
	const html = `
        <!DOCTYPE html>
        <html>
            <head>
                <meta charset="UTF-8">
                <title>{{ .Title }}</title>
            </head>
            <body>
			<center>
                <h1>{{ .Heading }}</h1>
				<br>
                <img src="data:image/jpg;base64,{{ .ImageX }}" />
				<p>{{ .NameImageX }}</p>
				<hr>
				<img src="data:image/jpg;base64,{{ .ImageZ }}" />
				<p>{{ .NameImageZ }}</p>
				<br>
				<footer>
				<p>{{ .Subject }}</p>
				<p>{{ .Names }}</p>
				<p>{{ .Date }}</p>
				</footer>
			</center>
            </body>
        </html>
    `

	// Compilar la plantilla
	tmpl, err := template.New("html").Parse(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Definir los datos a insertar en la plantilla
	data := struct {
		Title      string
		Heading    string
		ImageX     string
		NameImageX string
		ImageZ     string
		NameImageZ string
		Subject    string
		Names      string
		Date       string
	}{
		Title:      "Web",
		Heading:    obtenerHost(),
		ImageX:     codificar64(),
		NameImageX: "",
		ImageZ:     codificar64(),
		NameImageZ: "",
		Subject:    "Computacion en la nube",
		Names:      "Samuel Bermudez, Cristian Barerra",
		Date:       fecha(),
	}

	// Aplicar la plantilla con los datos y enviar la respuesta al cliente
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fecha() string {
	t := time.Now()
	ahora := fmt.Sprintf("%d / %d / %d", t.Day(), t.Month(), t.Year())
	return ahora
}

func listaImagenes() []string {
	direccion := "D:\\Users\\Cristian\\Documents\\2023-1\\Computacion Nube\\Pruebas"

	var nombres []string
	var i int

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
	return nombres
}

func aleatorioImagen(n int, arreglo []string) string {
	aleatorio := rand.Intn(n)
	posicion := arreglo[aleatorio]
	return posicion
}

func codificar64() string {
	direccion := "D:\\Users\\Cristian\\Documents\\2023-1\\Computacion Nube\\Pruebas"

	var arreglo []string

	arreglo = listaImagenes()
	tamano := len(arreglo)
	archivoAleatorio := aleatorioImagen(tamano, arreglo)
	imagen := direccion + "\\" + archivoAleatorio

	byte, err := ioutil.ReadFile(imagen)
	if err != nil {
		log.Fatal(err)
	}

	img64 := base64.StdEncoding.EncodeToString(byte)
	return img64
}

func obtenerNombre(nombre string) string {
	return nombre
}

func obtenerHost() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return hostname
}
