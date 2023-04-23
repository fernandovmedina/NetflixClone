@author Fernando Alejandro Vazquez Medina

# Lenguajes usados para el proyecto
- Go
- HTML
- CSS
- JavaScript

# Frameworks usados para el proyecto
- Fiber (Go)
- React (JavaScript)

# Herramientas usadas para el proyecto
- Vite (JavaScript)

# Endpoint de la api rest de NetflixClone
- Get
    - /api/series => Reetorna todas las series
    - /api/series/:serie/:temporada/:capitulo => Reetorna la informacion de cierto capitulo
    - /api/series/:serie/:temporada/:capitulo/img => Reetorna el wallper de cierto capitulo
    - /api/series/:serie/img => Reetorna el wallper de alguna serie
    - /api/series/temporadas/:serie => Reetorna el numero de temporadas de algun serie
    - /api/series/capitulos/:serie/temporadas/:temporada => Reetorna el numero de capitulos de alguna temporada de cierta serie
    - /api/series/:serie/:temporada => Reetorna todos los capitulos de algun temporada de cierta serie
