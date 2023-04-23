package main

import (
	"log"
	"os"

	"github.com/fernandovmedina/NetflixClone/backend/src/config"
	"github.com/fernandovmedina/NetflixClone/backend/src/database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Declaramos una variable que guarde los errores al momento de ejecucion
	var err error
	// Abrimos el archivo .env
	if err = godotenv.Load(); err != nil {
		log.Printf("%s \n", "Ocurrio algun problema al momento de abrir el archivo .env")
	}
	// Guardamos las variables del archivo .env
	var (
		appPort string = os.Getenv("APP_PORT")
	)
	// Iniciamos la conexion a la base de datos
	if _, err = sql.OpenDB(); err != nil {
		panic(err.Error())
	}
	// Creamos una instancia de fiber
	var app *fiber.App = fiber.New(fiber.Config{
		StrictRouting: true,
	})
	// Middleware
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	// Inicializamos los endpoints de la api
	config.SetupApp(app)
	// Corremos la app
	log.Fatal(app.Listen(appPort))
}
