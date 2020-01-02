package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/knuckerr/go_rest/api/models"
	"github.com/knuckerr/go_rest/api/responses"
	"github.com/knuckerr/go_rest/api/validators"
	"net/http"
)

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusOK, &users)

}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	err = validators.New(user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	user_created, err := user.SaveUser(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, user_created)

}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	user_id := chi.URLParam(r, "id")
	err := user.DeleteUser(server.DB, user_id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, map[string]string{"User id:" + user_id: "deleted"})
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	user_id := chi.URLParam(r, "id")
	find_user, err := user.FindUser(server.DB, user_id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, find_user)
}

//todo check if the user id is equal with the login user id
func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	err = validators.New(user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	user_id := chi.URLParam(r, "id")
	update_user, err := user.UpdateUser(server.DB, user_id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, update_user)
}
