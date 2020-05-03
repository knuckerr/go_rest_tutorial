package conf

import (
	"github.com/spf13/viper"
	"os"
)

func default_init() {
	// Server viperuration
	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", "9000")

	// Database Settings
	viper.SetDefault("storage.host", os.Getenv("DATABASE_HOST"))
	viper.SetDefault("storage.port", "5432")
	viper.SetDefault("storage.database", "apidb")
	viper.SetDefault("storage.username", "postgres")
	viper.SetDefault("storage.password", "1234")

	// Secret Key
	viper.SetDefault("SECRET_KEY", "asjdlasjldasjj21lj1233213@*&@(*&(@))")
	// REDIS
	viper.SetDefault("redis.host", os.Getenv("REDIS_HOST"))
	viper.SetDefault("redis.port", "6379")

}
