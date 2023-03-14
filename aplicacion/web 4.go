package main

import (
	"fmt"
	"net/http"
)

func main() {

	const P = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Prueba golang</title>
	</head>
	<body>
		<div style="text-align:center;">
		<h1>Nombre maquina</h1>
  		</div>
	</body>
	</html>
	`

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, P)
	})

	// Inicia el servidor en el puerto 8090
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
