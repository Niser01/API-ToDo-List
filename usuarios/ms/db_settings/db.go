package dbsettings

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_database := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_sslmode := os.Getenv("SSL_MODE")

	db_connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		db_host, db_user, db_password, db_database, db_port, db_sslmode,
	)

	DB, err = gorm.Open(postgres.Open(db_connection), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	fmt.Println("Conexión con la Base de Datos establecida correctamente.")
}
