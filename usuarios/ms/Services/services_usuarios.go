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

// ****************************************************************
// Modelo: Tipo_de_perfiles
type CreateTiposDePerfilesInput struct {
	Tipo_de_perfil string
	Nombre_animal  string
	Personalidad   string
	Frase_clave    string
}

func CreateTipo_de_perfiles(input CreateTiposDePerfilesInput) error {
	creacion := Model.Tipo_de_perfiles{
		Tipo_de_perfil: input.Tipo_de_perfil,
		Nombre_animal:  input.Nombre_animal,
		Personalidad:   input.Personalidad,
		Frase_clave:    input.Frase_clave,
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
	Id             int     `json:"id"`
	Tipo_de_perfil *string `json:"tipo_de_perfil,omitempty"`
	Nombre_animal  *string `json:"nombre_animal,omitempty"`
	Personalidad   *string `json:"personalidad,omitempty"`
	Frase_clave    *string `json:"frase_clave,omitempty"`
}

func UpdateTipo_de_perfiles(input UpdateTipo_de_perfilesInput) error {
	var registro_a_actualizar Model.Tipo_de_perfiles
	err := DataBase.DB.First(&registro_a_actualizar, input.Id).Error
	if err != nil {
		return err
	}
	//De acuerdo con los datos que se queran actualizar se escoje que se actualiza y que no
	//permitiendo la actualizacion de 1 o varios campos con un solo metodo
	if input.Tipo_de_perfil != nil {
		registro_a_actualizar.Tipo_de_perfil = *input.Tipo_de_perfil
	}
	if input.Nombre_animal != nil {
		registro_a_actualizar.Nombre_animal = *input.Nombre_animal
	}
	if input.Personalidad != nil {
		registro_a_actualizar.Personalidad = *input.Personalidad
	}
	if input.Frase_clave != nil {
		registro_a_actualizar.Frase_clave = *input.Frase_clave
	}

	err = DataBase.DB.Save(&registro_a_actualizar).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteTipo_de_perfiles(id int) error {
	// Se crea una instancia con solo el ID del registro a eliminar
	eliminacion := Model.Tipo_de_perfiles{ID: id}
	//Se realiza la eliminacion del registro pasando la direccion del registro
	err := DataBase.DB.Delete(&eliminacion).Error
	if err != nil {
		return err
	}
	return nil
}
