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

func (c *coleccion) addPelicula(id, anio int, nom, dir string) {
	p := c.crearPelicula(id, anio, nom, dir)
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

func (c *coleccion) crearPelicula(id, anio int, nom, dir string) pelicula {
	p := pelicula{id, nom, dir, anio}
	return p
}

func (c *coleccion) actualizarPelicula(id, anio int, nom, dir string) error {
	i, err := c.buscarPeliculaID(id)
	if err == nil {
		c.peliculas[i].nombre = nom
		c.peliculas[i].anio = anio
		c.peliculas[i].director = dir
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
	c.addPelicula(1, 1990, "Rambo", "Rambo")
	c.addPelicula(2, 1990, "Rambo 2", "Rambo")
	c.addPelicula(3, 1990, "Rambo 3", "Rambo")
	c.imprimirColeccion()
	c.actualizarPelicula(2, 1880, "Rambo 0", "Rambo")
	c.imprimirColeccion()

}
