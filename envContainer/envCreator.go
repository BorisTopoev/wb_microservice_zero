package envContainer

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvContainer struct {
	DatabaseURL string
	ClusterURLS string
	ClusterID   string
	Subject     string
}

func Get_env() EnvContainer {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	envContainer := EnvContainer{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		ClusterURLS: os.Getenv("ClusterURLS"),
		ClusterID:   os.Getenv("ClusterID"),
		Subject:     os.Getenv("Subject")}
	return envContainer
}
