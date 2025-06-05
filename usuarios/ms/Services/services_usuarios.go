package Services

import (
	"usuarios/ms/DataBase"
	"usuarios/ms/Model"
)

//Son los metodos que interactuan con gorm para hacer las peticiones CRUD a la BD

// Modelo: TipoDeAutenticacion
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
	return tipos, nil
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

// Modelo: Tipo_de_perfiles
type CreateTiposDePerfilesInput struct {
	tipo_perfil   string
	nombre_animal string
	personalidad  string
	frase_clave   string
}

func CreateTipo_de_perfiles(input CreateTiposDePerfilesInput) error {
	creacion := Model.Tipo_de_perfiles{
		Tipo_de_perfil: input.tipo_perfil,
		Nombre_animal:  input.nombre_animal,
		Personalidad:   input.personalidad,
		Frase_clave:    input.frase_clave,
	}
	err := DataBase.DB.Create(&creacion).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadFullTipo_de_perfiles() ([]Model.Tipo_de_perfiles, error) {
	var tipo_de_perfiles []Model.Tipo_de_perfiles
	err := DataBase.DB.Find(&tipo_de_perfiles).Error
	if err != nil {
		return []Model.Tipo_de_perfiles{}, err
	}
	return tipo_de_perfiles, nil
}

func ReadTipo_de_perfilesById(id int) (Model.Tipo_de_perfiles, error) {
	var tipo_de_perfiles Model.Tipo_de_perfiles
	err := DataBase.DB.First(&tipo_de_perfiles, id).Error
	if err != nil {
		return Model.Tipo_de_perfiles{}, err
	}
	return tipo_de_perfiles, nil
}

type UpdateTipo_de_perfilesInput struct {
	id            int
	tipo_de_peril string
	nombre_animal string
	personalidad  string
	frase_clave   string
}

func UpdateTipo_de_perfiles(input UpdateTipo_de_perfilesInput) error {
	var registro_a_actualizar Model.Tipo_de_perfiles
	err := DataBase.DB.First(&registro_a_actualizar, input.id)
	if err != nil {
		return err.Error
	}
	registro_a_actualizar.Tipo_de_perfil = input.tipo_de_peril
	registro_a_actualizar.Nombre_animal = input.nombre_animal
	registro_a_actualizar.Personalidad = input.personalidad
	registro_a_actualizar.Frase_clave = input.frase_clave
	err = DataBase.DB.Save(&registro_a_actualizar)
	if err != nil {
		return err.Error
	}
	return nil
}
