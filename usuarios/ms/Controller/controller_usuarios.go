package Controller

import (
	"encoding/json"
	"net/http"
	"usuarios/ms/Services"
)

func ControllerCreateTipoDeAutenticacion(w http.ResponseWriter, r *http.Request) {
	var body struct {
		NuevoTipo string `json:"NuevoTipo"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}
	err = Services.CreateTipoDeAutenticacion(body.NuevoTipo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
