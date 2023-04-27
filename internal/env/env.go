package env

import (
	"fmt"

	"github.com/spf13/viper"
)

func ViperEnvVariable(key string) string {
	// viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		fmt.Println("Invalid type assertion")
	}
	return value
}
