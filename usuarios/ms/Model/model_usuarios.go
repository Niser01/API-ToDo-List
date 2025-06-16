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

type Autenticaciones struct {
	ID                   int       `gorm:"primaryKey;column:id" json:"id"`
	Correo               string    `gorm:"column:correo" json:"correo"`
	Password_hash        string    `gorm:"column:password_hash" json:"password_hash"`
	Verificado           bool      `gorm:"column:verificado" json:"verificado"`
	Oauth_uid            string    `gorm:"column:oauth_uid" json:"oauth_uid"`
	Tipo_autenticacionID int       `gorm:"column:tipo_autenticacion" json:"tipo_autenticacion"`
	Created_at           time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	Updated_at           time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

type Usuarios struct {
	ID               int       `gorm:"primaryKey;column:id" json:"id"`
	AutenticacionID  int       `gorm:"column:autenticacion_id" json:"autenticacion_id"`
	Nombre_preferido string    `gorm:"column:nombre_preferido" json:"nombre_preferido"`
	Perfil_activo    bool      `gorm:"column:perfil_activo" json:"perfil_activo"`
	TipoPerfilID     int       `gorm:"column:tipo_perfil" json:"tipo_perfil"`
	Created_at       time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	Updated_at       time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
