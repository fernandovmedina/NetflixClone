package img

import (
	"net/http"

	database "github.com/fernandovmedina/NetflixClone/backend/src/database/sql"
	functions "github.com/fernandovmedina/NetflixClone/backend/src/functions"
	fiber "github.com/gofiber/fiber/v2"
)

// Endpoint para obtener la ruta de una imagen
// http://127.0.0.1:8080/api/series/nombre_serie/numero_temporada/numero_capitulo/img
func GetCapituloImg(c *fiber.Ctx) error {
	// Declaramos una variable que contenga la ruta del archivo
	var route string
	// Obtenemos los params
	var serie string = functions.CleanParams(c.Params("serie"))
	var temporada string = functions.CleanParams(c.Params("temporada"))
	var capitulo string = functions.CleanParams(c.Params("capitulo"))
	// Ejecutamos el query
	rows, err := database.DB.Query("select capitulos.img_capitulo_url from series, temporadas, capitulos where series.id_serie = temporadas.id_serie and temporadas.id_temporada = capitulos.id_temporada and series.nombre_serie = ? and temporadas.numero_temporada = ? and capitulos.numero_capitulo = ?", &serie, &temporada, &capitulo)
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
		// Escaneamos la route de la base de datos
		err = rows.Scan(&route)
		// Verificamos que no ocurra algun error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err,
				"data":    nil,
			})
		}
	}
	// Reetornamos el archivo
	return c.SendFile(route)
}

// Endpoint para obtener el wallper de una serie
func GetSerieWallper(c *fiber.Ctx) error {
	// Declaramos una variable que guarde la ruta de la imagen
	var route string
	// Obtenemos la serie mediante params
	var serie string = functions.CleanParams(c.Params("serie"))
	// Preparamos el query
	rows, err := database.DB.Query("select img_serie_url from NetflixDatabase.series where nombre_serie = ?", &serie)
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
		// Escaneamos la route de la imagen
		err = rows.Scan(&route)
		// Verificamos que no ocurra algun error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err,
				"data":    nil,
			})
		}
	}
	// Reetornamos la imagen
	return c.SendFile(route)
}
