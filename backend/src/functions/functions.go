package functions

import (
	"os"
	"strings"

	database "github.com/fernandovmedina/NetflixClone/backend/src/database/sql"
	"github.com/joho/godotenv"
)

// Funcion para limpiar los params
func CleanParams(params string) string {
	// Verificamos que los params contengan el signo de "-"
	if strings.Contains(params, "-") {
		return strings.ReplaceAll(params, "-", " ")
	}
	// Reetornamos los params tal y como estan
	return params
}

// Funcion para contar el numero de series en la base de datos
func CountSeries() int {
	// Declaramos una variable que guarde el count de la base de datos
	var count int
	// Ejecutamos el query para contar el numero de series
	stmt, err := database.DB.Query("select count(*) from NetflixDatabase.series")
	// Verificamos que no ocurra algun error
	if err != nil {
		return 0
	}
	// Hacemos que stmt sea lo ultimo en ejecutarse
	defer stmt.Close()
	// Recorremos las rows generadas por el query
	for stmt.Next() {
		// Escaneamos el count
		err = stmt.Scan(&count)
		// Verificamos que no ocurra algun error
		if err != nil {
			return 0
		}
	}
	// Reetornamos el count
	return count
}

// Funcion para obtener la llave secreta
func GetSecretKey() string {
	// Cargamos el archivo .env
	err := godotenv.Load()
	// Verificamos que no ocurra algun error
	if err != nil {
		return ""
	}
	// Retornamos la secret key desde el archivo .env
	return os.Getenv("SECRET_KEY")
}
