package config

import (
	"log"
	"github.com/spf13/viper"
)

type Configurations struct {
	Server ServerConfigurations
	PostgreSQL PostgreSQLConfigurations
	Redis RedisConfigurations
}

type ServerConfigurations struct {
	Addr string
}

type PostgreSQLConfigurations struct {
	Host			string
	User			string
	Password		string
	Port			int
	Database		string
	SSLMode			string
	Timezone		string
}

type RedisConfigurations struct {
	Auth string
}

func Config () *Configurations {

	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	var config Configurations

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return &config
}
