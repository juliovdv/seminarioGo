package main

import (
	"errors"
	"fmt"
)

type pelicula struct {
	id       int
	nombre   string
	director string
	anio     int
}

type coleccion struct {
	peliculas []pelicula
}

func (c *coleccion) addPelicula(p pelicula) {
	c.peliculas = append(c.peliculas, p)
}

func (c *coleccion) deletePeliculaPosicion(i int) {
	c.peliculas = c.peliculas[:i+copy(c.peliculas[i:], c.peliculas[i+1:])] //a = a[:i+copy(a[i:], a[i+1:])]
}

func (c *coleccion) buscarPeliculaID(id int) (int, error) {
	for i := 0; i < len(c.peliculas); i++ {
		if c.peliculas[i].id == id {
			return i, nil
		}
	}
	return -1, errors.New("este id no existe")
}

func (c *coleccion) deletePelicula(id int) error {
	i, err := c.buscarPeliculaID(id)
	if err == nil {
		c.deletePeliculaPosicion(i)

	}
	return err
}

func (c *coleccion) actualizarPelicula(p pelicula, id int) error {
	i, err := c.buscarPeliculaID(id)
	if err == nil {
		c.peliculas[i].nombre = p.nombre
		c.peliculas[i].anio = p.anio
		c.peliculas[i].director = p.director
	}
	return err
}

func (c *coleccion) getPeliculaID(id int) (pelicula, error) {
	i, err := c.buscarPeliculaID(id)
	return c.peliculas[i], err
}

func (c *coleccion) getPeliculas() []pelicula {
	return c.peliculas
}

func (c *coleccion) imprimirColeccion() {
	for i, p := range c.peliculas {
		fmt.Println(i, p)
	}
}

func main() {

	c := coleccion{}
	p1 := pelicula{1, "Rambo", "Rambo", 1990}
	p2 := pelicula{2, "Rambo 2", "Rambo", 1991}
	p3 := pelicula{3, "Rambo 3", "Rambo", 1990}
	c.addPelicula(p1)
	c.addPelicula(p2)
	c.addPelicula(p3)
	c.imprimirColeccion()
	c.actualizarPelicula(p3, 2)
	c.imprimirColeccion()

}
