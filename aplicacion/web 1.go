package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Manejador para todas las solicitudes entrantes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hola mundo")
	})

	// Crea un enrutador HTTP utilizando el paquete "net/http"
	router := http.NewServeMux()

	// Asigna el manejador a la ruta raíz
	router.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	}))

	// Inicia el servidor en el puerto 8085
	err := http.ListenAndServe(":8084", router)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
