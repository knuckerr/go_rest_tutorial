package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/knuckerr/go_rest/api/auth"
	"github.com/knuckerr/go_rest/api/models"
	"github.com/knuckerr/go_rest/api/responses"
	"net/http"
	"strings"
)

func (server *Server) Login(c *gin.Context) {
	login_user := models.LoginUser{}
	err := c.BindJSON(&login_user)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}
	db_user, err := login_user.Login(server.DB)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, errors.New("invaild crendetials"))
		return
	}
	token, err := auth.Createtoken(db_user)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(c, http.StatusOK, gin.H{"token": token})

}

func (Server *Server) RefreshToken(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	reqToken = strings.TrimSpace(splitToken[1])
	new_token, err := auth.Refreshtoken(reqToken)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(c, http.StatusOK, gin.H{"token": new_token})
}
