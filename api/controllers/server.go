package controllers

import (
	//"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	//"github.com/go-redis/redis/v7"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/knuckerr/go_rest/api/conf"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type Server struct {
	DB     *sqlx.DB
	Router *chi.Mux
	//Cache  *redis.Client
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
	server.DB, err = sqlx.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("error:", err)
	} else {
		log.Printf("We are connected to the %s database\n", Dbdriver)
	}
	/*
		client := redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		_, err = client.Ping().Result()
		if err != nil {
			log.Fatal("error:", err)

		}
		server.Cache = client
	*/
	driver, err := postgres.WithInstance(server.DB.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://api/migrations",
		"postgres", driver)
	err = m.Up()
	if err != nil {
		if err.Error() != "no change" {
			log.Fatal(err)
		}
	}
	server.Router = chi.NewRouter()
	server.Router.Use(middleware.RequestID)
	server.Router.Use(middleware.RealIP)
	server.Router.Use(middleware.Logger)
	server.Router.Use(middleware.Recoverer)
	server.InitializeRoutes()
}

func (server *Server) Run() {
	server.Initialize()
	log.Printf("Starting the server %s on port %s: ", viper.GetString("server.host"), viper.GetString("server.port"))
	http.ListenAndServe(":3000", server.Router)
}

/*
func (server *Server) Hset(key, field string, data interface{}) error {
	data_json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = server.Cache.HSet(key, field, data_json).Result()
	if err != nil {
		return err
	}
	return nil
}

func (server *Server) Hget(key, field string, data interface{}) error {
	v, err := server.Cache.HGet(key, field).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(v), data)
}
*/
