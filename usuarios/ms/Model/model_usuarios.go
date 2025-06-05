package Model

import "time"

//Son la representacion de las tablas de la bd en Go

type TipoDeAutenticacion struct {
	ID                int    `gorm:"primaryKey; column:id" json:"id"`
	TipoAutenticacion string `gorm:"column:tipo_de_autenticacion" json:"tipo_autenticacion"`
}

type Tipo_de_perfiles struct {
	ID             int    `gorm:"primaryKey;column:id" json:"id"`
	Tipo_de_perfil string `gorm:"column:tipo_de_perfil" json:"tipo_de_perfil"`
	Nombre_animal  string `gorm:"column:nombre_animal" json:"nombre_animal"`
	Personalidad   string `gorm:"column:personalidad" json:"personalidad"`
	Frase_clave    string `gorm:"column:frase_clave" json:"frase_clave"`
}

type Usuarios struct {
	ID               int       `gorm:"primaryKey;column:id" json:"id"`
	Nombre           string    `gorm:"column:nombre" json:"nombre"`
	Apellido         string    `gorm:"column:apellido" json:"apellido"`
	Nombre_preferido string    `gorm:"column:nombre_preferido" json:"nombre_preferido"`
	Perfil_activo    bool      `gorm:"column:perfil_activo" json:"perfil_activo"`
	Url_foto_pertif  string    `gorm:"column:url_foto_perfil" json:"url_foto_perfil"`
	Telefono         string    `gorm:"column:telefono" json:"telefono"`
	TipoPerfilID     int       `gorm:"column:tipo_perfil" json:"tipo_perfil"`
	Created_at       time.Time `gorm:"column:created_at" json:"created_at"`
	Updated_at       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type Autenticaciones struct {
	ID                   int       `gorm:"primaryKey;column:id"  json:"id"`
	Usuario_idID         int       `gorm:"column:usuario_id"  json:"usuario_id"`
	Correo               string    `gorm:"column:correo"  json:"correo"`
	Password_hash        string    `gorm:"column:password_hash"  json:"password_hash"`
	Oauth_uid            string    `gorm:"column:oauth_uid"  json:"oauth_uid"`
	Tipo_autenticacionID int       `gorm:"column:tipo_autenticacion"  json:"tipo_autenticacion"`
	Created_at           time.Time `gorm:"column:created_at"  json:"created_at"`
	Updated_at           time.Time `gorm:"column:updated_at"  json:"updated_at"`
}
