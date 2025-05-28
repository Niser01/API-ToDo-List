package main

import (
	"log"
	"net/http"

	"usuarios/ms/db_settings"
)

func main() {
	err := db_settings.InitDB()
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	log.Println("Conexión con la base de datos establecida correctamente.")
	log.Println("MS de usuarios inicializado")
	log.Println("Servidor escuchando en el puerto 8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
