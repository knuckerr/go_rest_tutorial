package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/knuckerr/go_rest/api/models"
	"github.com/knuckerr/go_rest/api/responses"
)

func (server *Server) GetUsers(c *gin.Context) {
	user := models.User{}
	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
	}
	responses.JSON(c, http.StatusOK, &users)

}

func (server *Server) CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}
	user_created, err := user.SaveUser(server.DB)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(c, http.StatusCreated, user_created)

}

func (server *Server) DeleteUser(c *gin.Context) {
	user := models.User{}
	user_id := c.Params.ByName("id")
	err := user.DeleteUser(server.DB, user_id)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"User id:" + user_id: "deleted"})
}

func (server *Server) GetUser(c *gin.Context) {
	user := models.User{}
	user_id := c.Params.ByName("id")
	find_user, err := user.FindUser(server.DB, user_id)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, find_user)
}

//todo check if the user id is equal with the login user id
func (server *Server) UpdateUser(c *gin.Context) {
	user := models.User{}
	user_id := c.Params.ByName("id")
	err := c.BindJSON(&user)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}
	update_user, err := user.UpdateUser(server.DB, user_id)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(c, http.StatusOK, update_user)
}
