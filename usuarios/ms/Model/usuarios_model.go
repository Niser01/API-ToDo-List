package model

import "time"

type TipoDeAutenticacion struct {
	ID                int    `gorm:"primaryKey;column:id"`
	TipoAutenticacion string `gorm:"column:tipo_de_autenticacion"`
}

type Tipo_de_perfiles struct {
	ID             int    `gorm:"primaryKey;column:id"`
	Tipo_de_perfil string `gorm:"column:tipo_de_perfil"`
	Personalidad   string `gorm:"column:personalidad"`
	Frase_clave    string `gorm:"column:frase_clave"`
}

type Usuarios struct {
	ID              int       `gorm:"primaryKey;column:id"`
	Nombre          string    `gorm:"column:nombre"`
	Apellido        string    `gorm:"column:apellido"`
	Perfil_activo   bool      `gorm:"column:perfil_activo"`
	Url_foto_pertif string    `gorm:"column:url_foto_perfil"`
	Telefono        string    `gorm:"column:telefono"`
	TipoPerfilID    int       `gorm:"column:tipo_perfil"`
	Created_at      time.Time `gorm:"column:created_at"`
	Updated_at      time.Time `gorm:"column:updated_at"`
}

type Autenticaciones struct {
	ID                   int       `gorm:"primaryKey;column:id"`
	Usuario_idID         int       `gorm:"column:usuario_id"`
	Correo               string    `gorm:"column:correo"`
	Password_hash        string    `gorm:"column:password_hash"`
	Oauth_uid            string    `gorm:"column:oauth_uid"`
	Tipo_autenticacionID int       `gorm:"column:tipo_autenticacion"`
	Created_at           time.Time `gorm:"column:created_at"`
	Updated_at           time.Time `gorm:"column:updated_at"`
}
