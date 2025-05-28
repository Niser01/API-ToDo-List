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
    tipo_de_perfil TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS usuarios (
    id SERIAL PRIMARY KEY,
    nombre TEXT,
    apellido TEXT,
    perfil_activo BOOLEAN DEFAULT TRUE,
    url_foto_perfil TEXT,
    telefono TEXT,
    tipo_perfil INTEGER REFERENCES tipo_de_perfiles(id) ON DELETE SET NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS autenticaciones (
    id SERIAL PRIMARY KEY,
    usuario_id INTEGER UNIQUE REFERENCES usuarios(id) ON DELETE CASCADE,
    correo TEXT UNIQUE NOT NULL,
    password_hash TEXT, 
    oauth_uid TEXT,
    tipo_autenticacion INTEGER REFERENCES tipo_de_autenticacion(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO tipo_de_perfiles (tipo_de_perfil) 
VALUES ('Buho'), ('Zorro'), ('Gato'), ('Panda');

INSERT INTO tipo_de_autenticacion (tipo_de_autenticacion) 
VALUES ('Google'), ('Email');
