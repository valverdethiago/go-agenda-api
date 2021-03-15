package util

import "github.com/spf13/viper"

type Config struct {
	DBServer          string `mapstructure:"DB_SERVER"`
	DBName            string `mapstructure:"DB_NAME"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	PrometheusAddress string `mapstructure:"PROMETHEUS_SERVER"`
}

func LoadConfig(path string, env string) (Config, error) {
	var config Config
	viper.AddConfigPath(path + "/env")
	viper.SetConfigName(env)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}
