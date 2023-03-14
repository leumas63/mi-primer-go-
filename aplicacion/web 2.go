package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Definir el manejador HTTP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola, esta es una respuesta del servidor")
	})

	// Inicia el servidor en el puerto 8090
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
