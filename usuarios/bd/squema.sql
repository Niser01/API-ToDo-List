DROP TABLE IF EXISTS usuarios;

CREATE TABLE IF NOT EXISTS usuarios (
    id SERIAL PRIMARY KEY,
    usuario TEXT UNIQUE NOT NULL,
    nombre TEXT,
    apellido TEXT,
    correo TEXT UNIQUE NOT NULL,
    password_hash TEXT, 
    oauth_provider TEXT, 
    oauth_uid TEXT       
);