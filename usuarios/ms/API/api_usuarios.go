package API

import (
	"usuarios/ms/Controller"

	"github.com/gorilla/mux"
)

func RutasAPITipoDeAutenticacionUsuarios(r *mux.Router) {
	ruta_base := r.PathPrefix("/tipo_de_autenticacion_usuarios").Subrouter()

	ruta_base.HandleFunc("/create", Controller.ControllerCreateTipoDeAutenticacion).Methods("POST")
	ruta_base.HandleFunc("/read", Controller.ControllerReadFullTipoDeAutenticacion).Methods("GET")
	ruta_base.HandleFunc("/read_by_id", Controller.ControllerReadTipoDeAutenticacionByID).Methods("GET")
	ruta_base.HandleFunc("/update", Controller.ControllerUpdateTipoDeAutenticacion).Methods("PUT")
	ruta_base.HandleFunc("/delete", Controller.ControllerDeleteTipoDeAutenticacion).Methods("DELETE")
}
