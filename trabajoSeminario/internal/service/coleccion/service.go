package coleccion

import (
	"seminarioGo/trabajoSeminario/internal/config"

	"github.com/jmoiron/sqlx"
)

//...
type Pelicula struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Director string `json:"director"`
	Anio     int    `json:"anio"`
}

//...
type Service interface {
	AddPelicula(Pelicula) error
	BusquedaID(int) *Pelicula
	GetPeliculas() []*Pelicula
}

type service struct {
	db *sqlx.DB
	c  *config.Config
}

//...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddPelicula(p Pelicula) error {
	return nil
}
func (s service) BusquedaID(ID int) *Pelicula {
	return nil
}
func (s service) GetPeliculas() []*Pelicula {
	var lista []*Pelicula
	s.db.Select(&lista, "SELECT * FROM coleccion")
	return lista
}
