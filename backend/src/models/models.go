package models

// Creamos una struct para registro
type Registro struct {
	NombreSerie     string `json:"nombre_serie"`
	NumeroTemporada int    `json:"numero_temporada"`
	NumeroCapitulo  int    `json:"numero_capitulo"`
	NombreCapitulo  string `json:"nombre_capitulo"`
	Duracion        string `json:"duracion"`
}

// Creamos un struct para user
type User struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Struct para el registro de capitulos
type RegistroCapitulo struct {
	NumeroTemporada int    `json:"numero_temporada"`
	NumeroCapitulo  int    `json:"numero_capitulo"`
	NombreCapitulo  string `json:"nombre_capitulo"`
	Duracion        string `json:"duracion"`
}
