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
    personalidad TEXT NOT NULL,
    frase_clave TEXT NOT NULL
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

INSERT INTO tipo_de_perfiles (tipo_de_perfil, personalidad, frase_clave) 
VALUES ('Buho', 'Le encanta tener todo bajo control. Se siente cómodo con estructuras claras, listas bien organizadas y planes establecidos. No necesita motivación externa, solo herramientas que le ayuden a optimizar su productividad.', '“Lo que no se planea, no se cumple.”'), 
        ('Zorro', 'Le motiva ver resultados y progresos. Busca mejorar constantemente y valora mucho los logros. Tiende a trabajar por objetivos y le gustan los desafíos. Es ágil, enfocado y competitivo de forma saludable.', '“Ver cuánto he avanzado me da más ganas.”'), 
        ('Gato', 'Sensible al entorno y a su estado de ánimo. Tiene días buenos y días difíciles, pero quiere mejorar. Valora mucho el apoyo emocional y las palabras motivadoras. Le cuesta arrancar, pero una vez en marcha, avanza.', '“A veces solo necesito que alguien crea en mí.”'), 
        ('Panda', 'Tranquilo y reflexivo. No le gusta sentirse saturado ni presionado. Prefiere visualizar sus opciones y decidir con calma. Parece relajado, pero cumple lo necesario. Su estilo es relajado pero efectivo.', '“Menos presión, más claridad.”');

INSERT INTO tipo_de_autenticacion (tipo_de_autenticacion) 
VALUES ('Google'), ('Email');
