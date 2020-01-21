package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Session struct {
		Secret     string
		MaxAge     int    `mapstructure:"max_age"`
		CookieName string `mapstructure:"cookie_name"`
	} `mapstructure:"session"`
	Database struct {
		Url string
	}
}

var Cfg Config

func Configure() {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "dev"
	}

	for _, name := range []string{"config", "config_" + env} {
		readConfigFile(name)
		if err := viper.Unmarshal(&Cfg); err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
	}

	bindEnv()
	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	log.Println("configuration is done")
}

func readConfigFile(name string) {
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Error config file: %s\n", err)
		} else {
			log.Panicf("Fatal error config file: %s\n", err)
		}
	}
}

func bindEnv() {
	viper.BindEnv("database.url", "DATABASE_URL")
	viper.BindEnv("session.secret", "APP_SECRET")
}
