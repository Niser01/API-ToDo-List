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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
		return
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ControllerDeleteTipoDeAutenticacion(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Id int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}
	err = Services.DeleteTipoDeAutenticacion(body.Id)
	if err != nil {
		mensaje := fmt.Sprintf("Error al eliminar el registro con id: %v - %v", body.Id, err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

//****************************************************************
//Controladores para Tipo_de_perfiles

func ControllerCreateTipo_de_perfiles(w http.ResponseWriter, r *http.Request) {
	var body_input Services.CreateTiposDePerfilesInput
	err := json.NewDecoder(r.Body).Decode(&body_input)
	if err != nil {
		http.Error(w, "JSON Invalido", http.StatusBadRequest)
		return
	}

	err = Services.CreateTipo_de_perfiles(body_input)
	if err != nil {
		mensaje := fmt.Sprintf("Error al crear el nuevo tipo de pertil. - %v", err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func ControllerReadFullTipo_de_perfiles(w http.ResponseWriter, r *http.Request) {
	query, err := Services.ReadFullTipo_de_perfiles()
	if err != nil {
		mensaje := fmt.Sprintf("Error al buscar los tipos de pertil. - %v", err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(query)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta a JSON", http.StatusInternalServerError)
		return
	}
}

func ControllerReadTipo_de_perfilesById(w http.ResponseWriter, r *http.Request) {
	var body_input struct {
		Id int `json:"id"`
	}
	var err error
	err = json.NewDecoder(r.Body).Decode(&body_input)
	if err != nil {
		http.Error(w, "Formato JSON invalido", http.StatusBadRequest)
		return
	}

	query, err := Services.ReadTipo_de_perfilesById(body_input.Id)
	if err != nil {
		mensaje := fmt.Sprintf("Error al realizar la query. - %v", err)
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(query)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta a JSON", http.StatusInternalServerError)
		return
	}
}

func ControllerUpdateTipo_de_perfiles(w http.ResponseWriter, r *http.Request) {
	var body_input Services.UpdateTipo_de_perfilesInput
	err := json.NewDecoder(r.Body).Decode(&body_input)
	if err != nil {
		http.Error(w, "Formato JSON invalido", http.StatusBadRequest)
	}
	err = Services.UpdateTipo_de_perfiles(body_input)
	if err != nil {
		message := fmt.Sprintf("Error al actualizar el perfil %v. - %v", body_input.Id, err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ControllerDeleteTipo_de_perfiles(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Id int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Formato de JSON invalido", http.StatusBadRequest)
		return
	}
	err = Services.DeleteTipo_de_perfiles(body.Id)
	if err != nil {
		mensaje := fmt.Sprintf("Error al eliminar el registro con id: %v. - %v", body.Id, err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

//********************************************************************
// Controladores para Autenticaciones

func ControllerCreateAutenticaciones(w http.ResponseWriter, r *http.Request) {
	var body_input Services.CreateAutenticacionesInput
	err := json.NewDecoder(r.Body).Decode(&body_input)
	if err != nil {
		http.Error(w, "JSON Invalido", http.StatusBadRequest)
		return
	}

	err = Services.CreateAutenticaciones(body_input)
	if err != nil {
		mensaje := fmt.Sprintf("Error al crear una nueva autenticacion. - %v,", err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func ControllerReadFullAutenticaciones(w http.ResponseWriter, r *http.Request) {
	query, err := Services.ReadFullAutenticaciones()
	if err != nil {
		mensaje := fmt.Sprintf("Error al buscar las autenticaciones - %v", err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(query)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta a JSON", http.StatusInternalServerError)
		return
	}
}

func ControllerReadAutenticacionesById(w http.ResponseWriter, r *http.Request) {
	var body_input struct {
		Id int `json:"id"`
	}
	var err error
	err = json.NewDecoder(r.Body).Decode(&body_input)
	if err != nil {
		http.Error(w, "Formato JSON inválido", http.StatusBadRequest)
		return
	}

	query, err := Services.ReadAutenticacionesById(body_input.Id)
	if err != nil {
		mensaje := fmt.Sprintf("Error al realizar la query. - %v", err)
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(query)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta a JSON", http.StatusInternalServerError)
		return
	}
}

func ControllerUpdateAutenticaciones(w http.ResponseWriter, r *http.Request) {
	var body_input Services.UpdateAutenticacionesInput
	err := json.NewDecoder(r.Body).Decode(&body_input)
	if err != nil {
		http.Error(w, "Formato JSON invalido", http.StatusBadRequest)
	}

	err = Services.UpdateAutenticaciones(body_input)
	if err != nil {
		mensaje := fmt.Sprintf("Error al actualizar la actualizacion %v. - %v", body_input.Id, err)
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ControllerDeleteAutenticaciones(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Id int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Formato JSON invalido", http.StatusBadRequest)
		return
	}
	err = Services.DeleteAutenticaciones(body.Id)
	if err != nil {
		mensaje := fmt.Sprintf("Error al eliminar la autenticacion con id: %v.  %v", body.Id, err.Error())
		http.Error(w, mensaje, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
