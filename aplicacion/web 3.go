package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html><head><title>[]</title></head><body><h1>[]</h1></body></html>")
	})

	// Inicia el servidor en el puerto 8090
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
