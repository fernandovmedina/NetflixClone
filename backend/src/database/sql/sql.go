package sql

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Creamos una variable que guarde la conexion a la base de datos
var DB *sql.DB

// Funcion para abrir la base de datos
func OpenDB() (*sql.DB, error) {
	// Declaramos una variable que guarde los errores al momento de ejecucion
	var err error
	// Abrimos el archivo .env
	if err = godotenv.Load(); err != nil {
		return nil, err
	}
	// Guardamos las variables del archivo .env
	var (
		databaseName     string = os.Getenv("DATABASE_NAME")
		databaseUser     string = os.Getenv("DATABASE_USER")
		databasePort     string = os.Getenv("DATABASE_PORT")
		databaseHost     string = os.Getenv("DATABASE_HOST")
		databasePassword string = os.Getenv("DATABASE_PASSWORD")
	)
	// Creamos el dsn para la base de datos
	var dsn string = databaseUser + ":" + databasePassword + "@tcp(" + databaseHost + ":" + databasePort + ")/" + databaseName
	// Nos conectamos a la base de datos
	if DB, err = sql.Open("mysql", dsn); err != nil {
		return nil, err
	} else {
		log.Println("Conexion exitosa a la base de datos")
		return DB, nil
	}
}
