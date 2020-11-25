package coleccion

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		path:     "/pelicula/:id",
		function: busquedaID(s),
	})
	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/pelicula/:id",
		function: borrarID(s),
	})
	return list

}

func getColeccion(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"coleccion": s.GetColeccion()})
	}
}

func busquedaID(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pelicula": s.BusquedaID(c.Param("id"))})
	}
}

func borrarID(s Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		s.BorrarID(c.Param("id"))
		c.JSON(http.StatusOK, nil)

	}
}
