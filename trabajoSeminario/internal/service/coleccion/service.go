package coleccion

import (
	"seminarioGo/trabajoSeminario/internal/config"

	"github.com/jmoiron/sqlx"
)

//...
type Pelicula struct {
	ID       int64  `json:"id"`
	Nombre   string `json:"nombre"`
	Director string `json:"director"`
	Anio     int    `json:"anio"`
}

//...
type Service interface {
	AddPelicula(Pelicula) *Pelicula
	BusquedaID(string) *Pelicula
	GetColeccion() []*Pelicula
	BorrarID(string) *Pelicula
	ModificarPelicula(string, Pelicula) *Pelicula
}

type service struct {
	db *sqlx.DB
	c  *config.Config
}

//...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddPelicula(p Pelicula) *Pelicula {
	insertarPelicula := `INSERT INTO coleccion (nombre, director, anio) VALUES ($1, $2, $3)`
	id, err := s.db.MustExec(insertarPelicula, p.Nombre, p.Director, p.Anio).LastInsertId()
	p.ID = id
	if err != nil {
		return nil
	}
	return &p
}
func (s service) BusquedaID(id string) *Pelicula {
	var p = Pelicula{}
	err := s.db.Get(&p, "SELECT * FROM coleccion WHERE id=$1", id)
	if err != nil {
		return nil
	}
	return &p
}
func (s service) GetColeccion() []*Pelicula {
	var lista []*Pelicula
	s.db.Select(&lista, "SELECT * FROM coleccion")
	return lista
}

func (s service) BorrarID(id string) *Pelicula {
	p := s.BusquedaID(id)
	_, err := s.db.MustExec("DELETE FROM coleccion WHERE id=$1", id).RowsAffected()
	if err != nil {
		return nil
	}
	return p
}

func (s service) ModificarPelicula(id string, p Pelicula) *Pelicula {
	modificarPelicula := "UPDATE coleccion SET nombre=$1, director=$2, anio=$3 WHERE id=$4"
	s.db.MustExec(modificarPelicula, p.Nombre, p.Director, p.Anio, id)
	pelicula := s.BusquedaID(id)
	return pelicula
}
