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
				<style>
		body {
			font-family: Calibri, sans-serif;
		}
		</style>
            </head>
            <body>
			<center>
                <h1>{{ .Heading }}</h1>
				<br>
                <img src="data:image/jpg;base64,{{ .ImageX }}" />
				<h4>{{ .NameImageX }}</h4>
				<hr>
				<img src="data:image/jpg;base64,{{ .ImageZ }}" />
				<h4>{{ .NameImageZ }}</h4>
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

	var imageX string
	var imageY string
	var nameX string
	var nameY string
	for {
		imageX, nameX = codificar64()
		imageY, nameY = codificar64()
		if nameX != nameY && imageX != imageY {
			break
		}
	}

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
		ImageX:     imageX,
		NameImageX: nameX,
		ImageZ:     imageY,
		NameImageZ: nameY,
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

func listaImagenes(direccion string) []string {

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

func codificar64() (string, string) {
	direccion := "D:\\Users\\Cristian\\Documents\\2023-1\\Computacion Nube\\Pruebas"

	var arreglo []string

	arreglo = listaImagenes(direccion)
	tamano := len(arreglo)
	archivoAleatorio := aleatorioImagen(tamano, arreglo)
	imagen := direccion + "\\" + archivoAleatorio

	byte, err := ioutil.ReadFile(imagen)
	if err != nil {
		log.Fatal(err)
	}

	img64 := base64.StdEncoding.EncodeToString(byte)
	return img64, archivoAleatorio
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

//https://www.digitalocean.com/community/tutorials/how-to-use-variables-and-constants-in-go-es
//https://steemit.com/cervantes/@orlmicron/array-y-slice-en-go-golang
//https://blastcoding.com/bucles-en-go-golang/
//https://www.tutorialesprogramacionya.com/goya/detalleconcepto.php?punto=21&codigo=21&inicio=15
//https://parzibyte.me/blog/2021/08/12/go-servir-carpeta-http/
//https://www.delftstack.com/es/howto/html/html-img-base64/
