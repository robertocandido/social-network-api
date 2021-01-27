package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// CreateUser insere um usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(body, &user); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare(); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepositories(db)
	user.ID, error = repository.Create(user)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

// UpdateUser atualiza um usuário no banco
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar Usuário!"))
}

// FindUser busta todos os utuários do banco
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuário!"))
}

// FindUsers busca um usuário no banco
func FindUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
	}
	defer db.Close()

	repository := repositories.NewUserRepositories(db)

	users, error := repository.Find(nameOrNick)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
	}

	response.JSON(w, http.StatusOK, users)
}

// DeleteUser remove um usuário do bando
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
