package Controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"usuarios/ms/Services"
)

//Son los metodos que interactuan con el API con los handlers REST

// Controladores para TipoDeAutenticacion
func ControllerCreateTipoDeAutenticacion(w http.ResponseWriter, r *http.Request) {
	//Se decodifica el json de la peticion HTTP y mapea el JSON en variables de Go
	var body struct {
		NuevoTipo string `json:"NuevoTipo"`
	}

	//Se le pasa como argumento al decode el puntero del body para que pueda acceder a modificarlo
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}
	err = Services.CreateTipoDeAutenticacion(body.NuevoTipo)
	if err != nil {
		mensaje := fmt.Sprintf("Error al Crear el nuevo tipo de autenticacion - %v", err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func ControllerReadFullTipoDeAutenticacion(w http.ResponseWriter, r *http.Request) {
	//En este metodo no se recibien vairables por lo que no hay que decodificar nada
	query, err := Services.ReadFullTipoDeAutenticacion()
	if err != nil {
		mensaje := fmt.Sprintf("Error al realizar la Query - %v", err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}
	//Pero si se codifica en formato JSON para entregar el resultado de la query
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(query)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta a JSON", http.StatusInternalServerError)
		return
	}
}

func ControllerReadTipoDeAutenticacionByID(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Id int `json:"id"`
	}

	//Se decodifica el json del body
	var err error
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Formato JSON invalido", http.StatusBadRequest)
	}

	//Se hace la peticion de la query
	query, err := Services.ReadTipoDeAutenticacionByID(body.Id)
	if err != nil {
		mensaje := fmt.Sprintf("Error al realizar la Query - %v", err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}

	//Se decodifica la respuesta del servidor
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(query)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta a JSON", http.StatusInternalServerError)
	}
}

func ControllerUpdateTipoDeAutenticacion(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Id             int    `json:"id"`
		TipoModificado string `json:"nuevo_tipo_autentiacion"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}
	err = Services.UpdateTipoDeAutenticacion(body.Id, body.TipoModificado)
	if err != nil {
		mensaje := fmt.Sprintf("Error al actualizar el registro con id: %v - %v", body.Id, err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}
}

func ControllerDeleteTipoDeAutenticacion(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Id int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusInternalServerError)
		return
	}
	err = Services.DeleteTipoDeAutenticacion(body.Id)
	if err != nil {
		mensaje := fmt.Sprintf("Error al eliminar el registro con id: %v - %v", body.Id, err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
	}
}

//****************************************************************
//Controladores para Tipo_de_perfiles

func ControllerCreateTipo_de_perfiles(w http.ResponseWriter, r *http.Request) {

}
