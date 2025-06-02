package Services

import (
	"usuarios/ms/DataBase"
	"usuarios/ms/Model"
)

func CrearTipoDeAutenticacion(nuevoTipo string) error {
	autenticacion := Model.TipoDeAutenticacion{TipoAutenticacion: nuevoTipo}
	err := DataBase.DB.Create(&autenticacion).Error
	if err != nil {
		return err
	}
	return nil
}
