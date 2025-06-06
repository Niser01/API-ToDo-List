DROP TABLE IF EXISTS tipo_de_autenticacion;
DROP TABLE IF EXISTS tipo_de_perfiles;
DROP TABLE IF EXISTS autenticaciones;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE IF NOT EXISTS tipo_de_autenticacion (
    id SERIAL PRIMARY KEY,
    tipo_de_autenticacion TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS tipo_de_perfiles(
    id SERIAL PRIMARY KEY,
    tipo_de_perfil TEXT UNIQUE NOT NULL,
    nombre_animal TEXT NOT NULL,
    personalidad TEXT NOT NULL,
    frase_clave TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS autenticaciones (
    id SERIAL PRIMARY KEY,
    correo TEXT UNIQUE NOT NULL,
    password_hash TEXT, 
    verificado BOOLEAN DEFAULT false,
    oauth_uid TEXT,
    tipo_autenticacion INTEGER REFERENCES tipo_de_autenticacion(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS usuarios (
    id SERIAL PRIMARY KEY,
    autenticacion_id INTEGER REFERENCES autenticaciones(id) ON DELETE CASCADE,
    nombre_preferido TEXT,
    perfil_activo BOOLEAN DEFAULT TRUE,
    tipo_perfil INTEGER REFERENCES tipo_de_perfiles(id) ON DELETE SET NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

