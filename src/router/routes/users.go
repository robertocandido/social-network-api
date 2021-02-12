package routes

import (
	"net/http"

	"api/src/controllers"
)

var userRoutes = []Route{
	{
		URI:                "/users",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindUsers,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/user/{id}/follow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.Follow,
		RequerAutenticacao: true,
	},
	{
		URI:                "/user/{id}/unfollow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.Unfollow,
		RequerAutenticacao: true,
	},
	{
		URI:                "/user/{id}/followers",
		Metodo:             http.MethodGet,
		Funcao:             controllers.Followers,
		RequerAutenticacao: true,
	},
}
