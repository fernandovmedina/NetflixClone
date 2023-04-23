package data

import (
	"net/http"
	"strconv"

	database "github.com/fernandovmedina/NetflixClone/backend/src/database/sql"

	"github.com/fernandovmedina/NetflixClone/backend/src/functions"
	"github.com/fernandovmedina/NetflixClone/backend/src/models"
	"github.com/gofiber/fiber/v2"
)

// Endpoint para obtener todas las series
// http://127.0.0.1:8080/api/series
func GetSeries(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	// Declaramos un mapa que guarde por llave el nombre de la serie y por valor un arreglo de registros
	var series map[string][]models.Registro = make(map[string][]models.Registro)
	// Obtenemos el numero de series de la base de datos
	var count int = functions.CountSeries()
	// Recorremos las series
	for i := 1; i <= count; i++ {
		// Ejecutamos el query
		rows, err := database.DB.Query("select series.nombre_serie, temporadas.numero_temporada, capitulos.numero_capitulo, capitulos.nombre_capitulo, capitulos.duracion from series, temporadas, capitulos where series.id_serie = temporadas.id_serie and temporadas.id_temporada = capitulos.id_temporada and series.id_serie = " + strconv.Itoa(i))
		// Verificamos que no ocurra algun error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err,
				"data":    nil,
			})
		}
		// Hacemos que rows sea lo ultimo en ejecutarse
		defer rows.Close()
		// Declaramos un arreglo de registros
		var registros []models.Registro
		// Recorremos las rows
		for rows.Next() {
			// Declaramos una variable de tipo registro
			var registro models.Registro
			// Escanemos los datos y los asignamos al registro
			err = rows.Scan(&registro.NombreSerie, &registro.NumeroTemporada, &registro.NumeroCapitulo, &registro.NombreCapitulo, &registro.Duracion)
			// Verificamos que no ocurra algun error
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"success": false,
					"message": err,
					"data":    nil,
				})
			}
			// Agregamos el registro al arreglo de registros
			registros = append(registros, registro)
		}
		// Ejecutamos un query para obtener el nombre de la serie en la que nos encontramos
		rows, err = database.DB.Query("select nombre_serie from NetflixDatabase.series where id_serie = " + strconv.Itoa(i))
		// Verificamos que no ocurra algun error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err,
				"data":    nil,
			})
		}
		// Declaramos una variable que contenga el nombre de la serie en la que nos encontramos
		var nombre string
		// Hacemos que rows sea lo ultimo en ejecutarse
		defer rows.Close()
		// Recorremos las rows
		for rows.Next() {
			// Escanemos el nombre
			err = rows.Scan(&nombre)
			// Verificamos que no ocurra algun error
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"success": false,
					"message": err,
					"data":    nil,
				})
			}
		}
		// Agregamos la serie al mapa de series
		series[nombre] = registros
	}
	// Reetornamos las series en formato json
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "series found",
		"data":    &series,
	})
}

// Endpoint para obtener una toda la informacion de un solo capitulo
// http://127.0.0.1:8080/api/series/:serie/:temporada/:capitulo
func GetSerie(c *fiber.Ctx) error {
	// Declaramos una variable de tipo registro
	var registro models.Registro
	// Obtenemos los params
	var serie string = functions.CleanParams(c.Params("serie"))
	var temporada string = functions.CleanParams(c.Params("temporada"))
	var capitulo string = functions.CleanParams(c.Params("capitulo"))
	// Ejecutamos el query para buscar el capitulo en la base de datos
	rows, err := database.DB.Query("select series.nombre_serie, temporadas.numero_temporada, capitulos.numero_capitulo, capitulos.nombre_capitulo, capitulos.duracion from series, temporadas, capitulos where series.id_serie = temporadas.id_serie and temporadas.id_temporada = capitulos.id_temporada and series.nombre_serie = ? and temporadas.numero_temporada = ? and capitulos.numero_capitulo = ?", &serie, &temporada, &capitulo)
	// Verificamos que no ocurra algun error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}
	// Hacemos que rows sea lo ultimo en ejecutarse
	defer rows.Close()
	// Recorremos las rows
	for rows.Next() {
		// Escaneamos los datos
		err = rows.Scan(&registro.NombreSerie, &registro.NumeroTemporada, &registro.NumeroCapitulo, &registro.NombreCapitulo, &registro.Duracion)
		// Verificamos que no ocurra algun error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err,
				"data":    nil,
			})
		}
	}
	// Reetornamos el registro en formato json
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "chapter found",
		"data":    &registro,
	})
}

// Endpoint para obtener el numero de temporadas de una serie
// http://127.0.0.1:8080/api/series/:serie/numero-temporadas
func GetNumeroTemporadas(c *fiber.Ctx) error {
	// Declaramos una variable que contenga el numero de temporadas
	var temporadas int
	// Obtenemos los params
	var serie string = functions.CleanParams(c.Params("serie"))
	// Ejecutamos el query para obtener el numero de temporadas de una serie
	rows, err := database.DB.Query("select count(temporadas.numero_temporada) from NetflixDatabase.series, NetflixDatabase.temporadas where series.id_serie = temporadas.id_serie and  series.nombre_serie = ?", &serie)
	// Verificamos que no ocurra algun error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}
	// Hacemos que rows sea lo ultimo en ejecutarse
	defer rows.Close()
	// Recorremos las rows
	for rows.Next() {
		// Escanemos el numero de temporadas
		err = rows.Scan(&temporadas)
		// Verificamos que no ocurra algun error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err,
				"data":    nil,
			})
		}
	}
	// Retornamos el numero de temporadas en formato json
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "number of seasons found",
		"data":    &temporadas,
	})
}

// Endpoint para obtener el numero de capitulos de una temporada
// http://127.0.0.1:8080/api/series/capitulos/:serie/:temporada
func GetNumeroCapitulos(c *fiber.Ctx) error {
	// Declaramos una variable que contenga el numero de capitulos de algun temporada
	var capitulos int
	// Obtenemos los params
	var serie string = functions.CleanParams(c.Params("serie"))
	var temporada string = functions.CleanParams(c.Params("temporada"))
	// Ejecutamos el query para obtener el numero de capitulos de alguna temporada
	rows, err := database.DB.Query("select count(capitulos.numero_capitulo) from NetflixDatabase.series, NetflixDatabase.temporadas, NetflixDatabase.capitulos where series.id_serie = temporadas.id_serie and temporadas.id_temporada = capitulos.id_temporada and series.nombre_serie = ? and temporadas.numero_temporada = ?", &serie, &temporada)
	// Verificamos que no ocurra algun error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}
	// Hacemos que rows sea lo ultimo en ejecutarse
	defer rows.Close()
	// Recorremos las rows
	for rows.Next() {
		// Escaneamos el numero de capitulos
		err = rows.Scan(&capitulos)
		// Verificamos que no ocurra algun error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err,
				"data":    nil,
			})
		}
	}
	// Reetornamos el numero de capitulos en formato json
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "number of chapters found",
		"data":    &capitulos,
	})
}

// Endpoint para obtener todos los capitulos de alguna temporada de alguna serie
// http://127.0.0.1:8080/api/series/:serie/:temporada
func GetCapitulos(c *fiber.Ctx) error {
	// Declaramos un arreglo de RegistroCapitulo
	var capitulos []models.RegistroCapitulo
	// Obtenemos los params
	var serie string = functions.CleanParams(c.Params("serie"))
	var temporada string = functions.CleanParams(c.Params("temporada"))
	// Ejecutamos el query
	rows, err := database.DB.Query("select temporadas.numero_temporada, capitulos.numero_capitulo, capitulos.nombre_capitulo, capitulos.duracion from series, temporadas, capitulos where series.id_serie = temporadas.id_serie and temporadas.id_temporada = capitulos.id_temporada and series.nombre_serie = ? and temporadas.numero_temporada = ?", &serie, &temporada)
	// Verificamos que no ocurra algun error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}
	// Hacemos que rows sea lo ultimo en ejecutarse
	defer rows.Close()
	// Recorremos las rows
	for rows.Next() {
		// Declaramos una variable de RegistroCapitulo
		var capitulo *models.RegistroCapitulo = new(models.RegistroCapitulo)
		// Escaneamos las rows
		err = rows.Scan(&capitulo.NumeroTemporada, &capitulo.NumeroCapitulo, &capitulo.NombreCapitulo, &capitulo.Duracion)
		// Verificamos que no ocurra algun error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err,
				"data":    nil,
			})
		}
		// Agregamos el capitulo al arreglo de capitulos
		capitulos = append(capitulos, *capitulo)
	}
	// Reetornamos los capitulos en formato json
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "chapters found",
		"data":    &capitulos,
	})
}
