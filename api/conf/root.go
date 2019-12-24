package conf

import (
	"github.com/spf13/viper"
)

func Init() {
	viper.SetTypeByDefaultValue(true)
	viper.SetConfigFile("settings.toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		default_init()
	}
}
