package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type pelicula struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Director string `json:"director"`
	Anio     int    `json:"anio"`
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
		if c.peliculas[i].ID == id {
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
		c.peliculas[i].Nombre = nom
		c.peliculas[i].Anio = anio
		c.peliculas[i].Director = dir
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

	r := gin.Default()
	r.GET("/pelicula/:id", getApi)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getApi(c *gin.Context) {
	col := coleccion{}
	col.addPelicula(1, 1990, "Rambo", "Rambo")
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	msg, err := col.getPeliculaID(i)
	if err == nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, msg)
}
