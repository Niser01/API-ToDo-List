package API

import (
	"usuarios/ms/Controller"

	"github.com/gorilla/mux"
)

func RutasAPITipoDeAutenticacionUsuarios(r *mux.Router) {
	ruta_base := r.PathPrefix("/tipo_de_autenticacion").Subrouter()

	ruta_base.HandleFunc("/create", Controller.ControllerCreateTipoDeAutenticacion).Methods("POST")
	ruta_base.HandleFunc("/read", Controller.ControllerReadFullTipoDeAutenticacion).Methods("GET")
	ruta_base.HandleFunc("/read_by_id", Controller.ControllerReadTipoDeAutenticacionByID).Methods("GET")
	ruta_base.HandleFunc("/update", Controller.ControllerUpdateTipoDeAutenticacion).Methods("PUT")
	ruta_base.HandleFunc("/delete", Controller.ControllerDeleteTipoDeAutenticacion).Methods("DELETE")
}

func RutasAPITipoDePerfilesUsuarios(r *mux.Router) {
	ruta_base := r.PathPrefix("/tipo_de_perfil").Subrouter()

	ruta_base.HandleFunc("/create", Controller.ControllerCreateTipo_de_perfiles).Methods("POST")
	ruta_base.HandleFunc("/read", Controller.ControllerReadFullTipo_de_perfiles).Methods("GET")
	ruta_base.HandleFunc("/read_by_id", Controller.ControllerReadTipo_de_perfilesById).Methods("GET")
	ruta_base.HandleFunc("/update", Controller.ControllerUpdateTipo_de_perfiles).Methods("PUT")
	ruta_base.HandleFunc("/delete", Controller.ControllerDeleteTipo_de_perfiles).Methods("DELETE")
}

func RutasAPIAutenticaciones(r *mux.Router) {
	ruta_base := r.PathPrefix("/autenticacion").Subrouter()

	ruta_base.HandleFunc("/create", Controller.ControllerCreateAutenticaciones).Methods("POST")
	ruta_base.HandleFunc("/read", Controller.ControllerReadFullAutenticaciones).Methods("GET")
	ruta_base.HandleFunc("/read_by_id", Controller.ControllerReadAutenticacionesById).Methods("GET")
	ruta_base.HandleFunc("/update", Controller.ControllerUpdateAutenticaciones).Methods("PUT")
	ruta_base.HandleFunc("/delete", Controller.ControllerDeleteAutenticaciones).Methods("DELETE")
}
