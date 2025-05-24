DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS tareas;

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

CREATE TABLE IF NOT EXISTS tareas (
    id SERIAL PRIMARY KEY,
    id_usuario INT REFERENCES usuarios(id) ON DELETE CASCADE,
    nombre_tarea TEXT NOT NULL,
    fecha_tarea DATE,
    completada BOOLEAN DEFAULT FALSE,
    activa BOOLEAN DEFAULT TRUE
);