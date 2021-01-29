package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route representa todas as rotas da aplicação
type Route struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurate coloca todas as rotas dentro do router
func Configurate(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Funcao).Methods(route.Metodo)
	}

	return r
}
