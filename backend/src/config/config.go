package config

import (
	api "github.com/fernandovmedina/NetflixClone/backend/src/controllers/api/data"
	img "github.com/fernandovmedina/NetflixClone/backend/src/controllers/api/img"
	fiber "github.com/gofiber/fiber/v2"
)

// Funcion para inicializar los endpoints de la api
func SetupApp(app *fiber.App) {
	// Endpoint para obtener todas las series
	app.Get("/api/series", api.GetSeries)
	// Endpoint para obtener la informacion de un capitulo
	app.Get("/api/series/:serie/:temporada/:capitulo", api.GetSerie)
	// Endpoint para obtener la imagen de un capitulo
	app.Get("/api/series/:serie/:temporada/:capitulo/img", img.GetCapituloImg)
	// Endpoint para obtener la imagen de una serie
	app.Get("/api/series/:serie/img", img.GetSerieWallper)
	// Endpoint para obtener el numero de temporadas de una serie
	app.Get("/api/series/temporadas/:serie", api.GetNumeroTemporadas)
	// Endpoint para obtener el numero de capitulos de una temporada
	app.Get("/api/series/capitulos/:serie/temporadas/:temporada", api.GetNumeroCapitulos)
	// Endpoint para obtener todos los capitulos de alguna temporada
	app.Get("/api/series/:serie/:temporada", api.GetCapitulos)
}
