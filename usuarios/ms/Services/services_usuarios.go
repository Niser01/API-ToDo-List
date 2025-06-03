package Services

import (
	"usuarios/ms/DataBase"
	"usuarios/ms/Model"
)

func CreateTipoDeAutenticacion(nuevoTipo string) error {
	autenticacion := Model.TipoDeAutenticacion{TipoAutenticacion: nuevoTipo}
	err := DataBase.DB.Create(&autenticacion).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadFullTipoDeAutenticacion() ([]Model.TipoDeAutenticacion, error) {
	var tipos []Model.TipoDeAutenticacion
	err := DataBase.DB.Find(&tipos).Error
	if err != nil {
		return []Model.TipoDeAutenticacion{}, err
	}
	return tipos, err
}

func ReadTipoDeAutenticacionByID(id int) (Model.TipoDeAutenticacion, error) {
	var autenticacion Model.TipoDeAutenticacion
	err := DataBase.DB.First(&autenticacion, id).Error
	if err != nil {
		return Model.TipoDeAutenticacion{}, err
	}
	return autenticacion, nil
}

func UpdateTipoDeAutenticacion(id int, tipoModificado string) error {
	var item_a_actualizar Model.TipoDeAutenticacion
	err := DataBase.DB.First(&item_a_actualizar, id).Error
	if err != nil {
		return err
	}
	item_a_actualizar.TipoAutenticacion = tipoModificado
	err = DataBase.DB.Save(&item_a_actualizar).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteTipoDeAutenticacion(id int) error {
	eliminacion := Model.TipoDeAutenticacion{ID: id}
	err := DataBase.DB.Delete(&eliminacion).Error
	if err != nil {
		return err
	}
	return nil
}
