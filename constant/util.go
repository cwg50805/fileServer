package constant

import (
	"log"

	"github.com/spf13/viper"
)

func ReadConfig(configPath string) {
	viper.SetConfigFile(configPath)
	viper.AddConfigPath(".")

	viper.SetDefault("PORT", ":4300")
	viper.BindEnv("PORT")
	viper.SetDefault("RUN_MODE", "release")
	viper.BindEnv("RUN_MODE")
	viper.SetDefault("READ_TIMEOUT", 180)
	viper.BindEnv("READ_TIMEOUT")
	viper.SetDefault("WRITE_TIMEOUT", 60)
	viper.BindEnv("WRITE_TIMEOUT")
	viper.SetDefault("REQUEST_TIMEOUT", 10000)
	viper.BindEnv("REQUEST_TIMEOUT")

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err.Error())
	}

}
