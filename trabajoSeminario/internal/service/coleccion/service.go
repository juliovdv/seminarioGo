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
	AddPelicula(Pelicula)
	BusquedaID(string) *Pelicula
	GetColeccion() []*Pelicula
	BorrarID(string)
}

type service struct {
	db *sqlx.DB
	c  *config.Config
}

//...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddPelicula(p Pelicula) {
	insertarPelicula := `INSERT INTO coleccion (nombre, director, anio) VALUES ($1, $2, $3)`
	s.db.MustExec(insertarPelicula, p.Nombre, p.Director, p.Anio)
}
func (s service) BusquedaID(id string) *Pelicula {
	var pelicula = Pelicula{}
	err := s.db.Get(&pelicula, "SELECT * FROM coleccion WHERE id=$1", id)
	if err != nil {
		return nil
	}
	return &pelicula
}
func (s service) GetColeccion() []*Pelicula {
	var lista []*Pelicula
	s.db.Select(&lista, "SELECT * FROM coleccion")
	return lista
}

func (s service) BorrarID(id string) {

	s.db.MustExec("DELETE FROM coleccion WHERE id=$1", id).RowsAffected()
}
