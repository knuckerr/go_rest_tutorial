package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/knuckerr/go_rest/api/conf"
	"github.com/knuckerr/go_rest/api/models"
	"github.com/spf13/viper"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

const Dbdriver = "postgres"

func (server *Server) Initialize() {
	conf.Init()
	var err error
	var db_host = viper.GetString("storage.host")
	var db_port = viper.GetString("storage.port")
	var db_user = viper.GetString("storage.username")
	var db_pass = viper.GetString("storage.password")
	var db_name = viper.GetString("storage.database")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", db_host, db_port, db_user, db_name, db_pass)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		log.Printf("We are connected to the %s database\n", Dbdriver)
	}
	server.DB.AutoMigrate(models.User{})
	server.Router = gin.Default()
	server.InitializeRoutes()
}

func (server *Server) Run() {
	server.Initialize()
	log.Printf("Starting the server %s on port %s: ", viper.GetString("server.host"), viper.GetString("server.port"))
	server.Router.Run()
}
