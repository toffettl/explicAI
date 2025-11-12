package configuration

import "github.com/spf13/viper"

var config = viper.New()

func Init() *viper.Viper {
	defaultConfigs()

	config.AutomaticEnv()
	return config
}

func defaultConfigs() {
	config.SetDefault("server:host", "0.0.0.0:8080")
	config.SetDefault("app.name", "explicaAI")
}
