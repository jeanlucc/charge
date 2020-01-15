package config

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Secret   string `yaml:"secret" envconfig:"APP_SECRET"`
	Database struct {
		Url string `yaml:"url" envconfig:"DATABASE_URL"`
	} `yaml:"database"`
}

var Cfg Config

func Configure() {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "dev"
	}

	readFile(&Cfg, "./config/config.yaml")
	readFile(&Cfg, "./config/config_"+env+".yaml")
	readEnv(&Cfg)
}

func readFile(cfg *Config, path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Println("Could not open config to initialise configuration: "+path, err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Println("Could not decode config to initialise configuration: "+path, err)
	}
}

func readEnv(cfg *Config) {
	toto := os.Getenv("DATABASE_URL")
	log.Println(toto)
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Panic("Could not initialise configuration with environment", err)
	}
}
