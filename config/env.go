package config

import "github.com/spf13/viper"

type EnvConf struct {
	Domain        string `mapstructure:"DOMAIN"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	Environment   string `mapstructure:"ENVIRONMENT"`
	SSLMode       string `mapstructure:"SSL_MODE"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	Version       string `mapstructure:"VERSION"`
}

func LoadEnv(path string) (*EnvConf, error) {
	var conf *EnvConf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return &EnvConf{}, err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return &EnvConf{}, err
	}

	return conf, nil
}
