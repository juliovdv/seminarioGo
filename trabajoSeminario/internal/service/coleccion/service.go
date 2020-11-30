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
	AddPelicula(Pelicula) (*Pelicula, error)
	BusquedaID(string) (*Pelicula, error)
	GetColeccion() ([]*Pelicula, error)
	BorrarID(string) (*Pelicula, error)
	ModificarPelicula(string, Pelicula) (*Pelicula, error)
}

type service struct {
	db *sqlx.DB
	c  *config.Config
}

//...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddPelicula(p Pelicula) (*Pelicula, error) {
	insertarPelicula := `INSERT INTO coleccion (nombre, director, anio) VALUES ($1, $2, $3)`
	id, err := s.db.MustExec(insertarPelicula, p.Nombre, p.Director, p.Anio).LastInsertId()
	p.ID = id
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func (s service) BusquedaID(id string) (*Pelicula, error) {
	var p = Pelicula{}
	err := s.db.Get(&p, "SELECT * FROM coleccion WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func (s service) GetColeccion() ([]*Pelicula, error) {
	var lista []*Pelicula
	err := s.db.Select(&lista, "SELECT * FROM coleccion")
	return lista, err
}

func (s service) BorrarID(id string) (*Pelicula, error) {
	p, err := s.BusquedaID(id)
	if err != nil {
		return nil, err
	}
	_, err = s.db.MustExec("DELETE FROM coleccion WHERE id=$1", id).RowsAffected()
	if err != nil {
		return nil, err
	}
	return p, err
}

func (s service) ModificarPelicula(id string, p Pelicula) (*Pelicula, error) {
	modificarPelicula := "UPDATE coleccion SET nombre=$1, director=$2, anio=$3 WHERE id=$4"
	s.db.MustExec(modificarPelicula, p.Nombre, p.Director, p.Anio, id)
	pelicula, err := s.BusquedaID(id)
	return pelicula, err
}
