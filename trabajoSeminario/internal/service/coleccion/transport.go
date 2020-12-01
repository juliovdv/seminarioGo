package coleccion

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

//...
func NewHTTPService(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/coleccion",
		function: getColeccion(s),
	})
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/coleccion/pelicula/:id",
		function: busquedaID(s),
	})
	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/coleccion/pelicula/:id",
		function: borrarID(s),
	})
	list = append(list, &endpoint{
		method:   "POST",
		path:     "/coleccion/pelicula",
		function: agregarPelicula(s),
	})
	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/coleccion/pelicula/:id",
		function: modificarPelicula(s),
	})
	return list

}

func getColeccion(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		rtn, err := s.GetColeccion()
		if err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"error":  "no se pudo obtener la coleccion",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"coleccion": rtn})
		}
	}
}

func busquedaID(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		rtn, err := s.BusquedaID(c.Param("id"))
		if err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"error":  "no se pudo obtener la pelicula",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"pelicula": rtn})
		}
	}
}

func borrarID(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		rtn, err := s.BorrarID(c.Param("id"))
		fmt.Println(err)
		if err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"error":  "no se pudo borrar la pelicula",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":   "borrado",
				"pelicula": rtn})
		}
	}
}

func agregarPelicula(s Service) gin.HandlerFunc {
	var p Pelicula
	return func(c *gin.Context) {
		rtn, err := s.AddPelicula(p)
		if err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"error":  "no se pudo agregar la pelicula",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":   "Agregado",
				"pelicula": rtn})
		}
	}
}

func modificarPelicula(s Service) gin.HandlerFunc {
	var p Pelicula

	return func(c *gin.Context) {
		rtn, err := s.ModificarPelicula(c.Param("id"), p)
		if err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"error":  "no se pudo modificar la pelicula"})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":   "Modificado",
				"pelicula": rtn})
		}
	}
}
