package controllers

import (
	"encoding/json"
	"errors"
	"github.com/knuckerr/go_rest/api/auth"
	"github.com/knuckerr/go_rest/api/models"
	"github.com/knuckerr/go_rest/api/responses"
	"net/http"
	"strings"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	login_user := models.LoginUser{}
	err := json.NewDecoder(r.Body).Decode(&login_user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	db_user, err := login_user.Login(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("invaild crendetials"))
		return
	}
	token, err := auth.Createtoken(db_user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, map[string]string{"token": token})

}

func (Server *Server) RefreshToken(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	reqToken = strings.TrimSpace(splitToken[1])
	new_token, err := auth.Refreshtoken(reqToken)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, map[string]string{"token": new_token})
}
