package auth

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fernandovmedina/NetflixClone/backend/src/database/sql"
	"github.com/fernandovmedina/NetflixClone/backend/src/functions"
	"github.com/fernandovmedina/NetflixClone/backend/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Funcion para registrar un nuevo usuario
func Register(c *fiber.Ctx) error {
	// Declaramos un mapa que guarde los datos
	var datos map[string]string = make(map[string]string)
	// Asignamos los datos mandados al mapa datos
	if err := c.BodyParser(&datos); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}
	// Encriptamos la contraseña
	password, err := bcrypt.GenerateFromPassword([]byte(datos["password"]), bcrypt.DefaultCost)
	// Verificamos que no ocurra ningun error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}
	// Declaramos una variable de tipo user y los guardamos en la base de datos
	var user models.User = models.User{
		Nombre:   datos["nombre"],
		Email:    datos["email"],
		Password: string(password),
	}
	// Ejecutamos un query para guardar el user en la base de datos
	_, err = sql.DB.Exec("insert into NetflixDatabase.users(nombre, email, password) values (?, ?, ?)", &user.Nombre, &user.Email, &user.Password)
	// Verificamos que no ocurra ningun error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}
	// Reetornamos el user en formato json
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "user saved",
		"data":    &user,
	})
}

// Funcion para el log-in
func Login(c *fiber.Ctx) error {
	// Declaramos un mapa que guarde los datos
	var datos map[string]string = make(map[string]string)
	// Obtenemos los datos mandados
	if err := c.BodyParser(&datos); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}
	// Declaramos una variable de tipo user
	var user models.User
	// Buscamos el user por medio del email
	rows, err := sql.DB.Query("select id, nombre, email, password from NetflixDatabase.users where email = ?", datos["email"])
	// Verificamos que no ocurra ningun error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}
	// Hacemos que rows sea lo ultimo en ejecutarse
	defer rows.Close()
	// Recorremos las rows
	for rows.Next() {
		// Escanemos las rows
		err = rows.Scan(&user.Id, &user.Nombre, &user.Email, &user.Password)
		// Verificamos que no ocurra ningun error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}
	// Verificamos que el user existe
	if user.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "user not found",
		})
	}
	// Comparamos la contraseña de la base de datos y la mandada por json
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(datos["password"])); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "wrong password",
		})
	}
	// Generamos las claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer: strconv.Itoa(user.Id),
	})
	// Generamos el token
	token, err := claims.SignedString([]byte(functions.GetSecretKey()))
	// Verificamos que no ocurra ningun error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not login",
		})
	}
	// Generamos la cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: true,
	}
	// Asignamos la cookie
	c.Cookie(&cookie)
	// Reetornamos un mensaje de exito
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// Funcion para obtener los datos de un user
func User(c *fiber.Ctx) error {
	// Obtenemos la cookie con value = "jwt"
	cookie := c.Cookies("jwt")
	// Desencriptamos los claims y obtenemos los datos
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(functions.GetSecretKey()), nil
	})
	// Verificamos que no ocurra ningun error
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	// Hacemos un casting
	claims := token.Claims.(*jwt.RegisteredClaims)
	// Declaramos una variable de tipo user
	var user models.User
	// Buscamos el user en la base de datos
	rows, err := sql.DB.Query("select id, nombre, email, password from NetflixDatabase.users where id = ?", claims.Issuer)
	// Verificamos que no ocurra ningun error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	// Hacemos que rows sea lo ultimo en ejecutarse
	defer rows.Close()
	// Recorremos las rows
	for rows.Next() {
		// Escaneamos las rows
		err = rows.Scan(&user.Id, &user.Nombre, &user.Email, &user.Password)
		// Verificamos que no ocurra ningun error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}
	// Reetornamos el user en formato json
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "user found",
		"data":    &user,
	})
}

// Funcion para cerrar la sesion
func Logout(c *fiber.Ctx) error {
	// Generamos una cookie y obtenemos los valores de la cookie actual
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	// Asignamos la cookie
	c.Cookie(&cookie)
	// Reetornamos un mensaje de exito
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
