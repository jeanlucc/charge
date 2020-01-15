package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnvPath(path string) {
	e := godotenv.Load(path)
	if e != nil {
		log.Println("Could not load environment file: "+path, e)
	}
}

func LoadEnv() {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "dev"
	}

	loadEnvPath(".env." + env + ".local")
	if "test" != env {
		loadEnvPath(".env.local")
	}
	loadEnvPath(".env." + env)
	loadEnvPath(".env")
}
