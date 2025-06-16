package main

import (
	"log"
	"net/http"

	"usuarios/ms/API"
	"usuarios/ms/DataBase"

	"github.com/gorilla/mux"
)

func main() {
	//Se genera la conexión con la base de datos
	err := DataBase.InitDB()
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}
	log.Println("Conexión con la base de datos establecida correctamente.")

	//Se incluyen las rutas del API
	route_handler := mux.NewRouter()
	API.RutasAPITipoDeAutenticacionUsuarios(route_handler)
	API.RutasAPITipoDePerfilesUsuarios(route_handler)
	API.RutasAPIAutenticaciones(route_handler)
	log.Println("MS de usuarios inicializado")

	//Se lanza el servidor en el puerto 8080
	log.Println("Servidor escuchando en el puerto 8080")
	err = http.ListenAndServe(":8080", route_handler)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
