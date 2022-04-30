package configs

import (
	"errors"
	"os"
)

type AppConfig struct {
	DbType string
}

const selectedDb = "mongo"

func NewAppConfig() {

}

func GetDbConfig() string {
	var dbUri string
	switch selectedDb {
	case "mongo":
		{
			mongoUri := os.Getenv("MONGO_URI")
			if mongoUri == "" {
				panic(errors.New("mongo uri not defined in env"))
			}
			dbUri = mongoUri
		}
	default:
		dbUri = ""
	}

	return dbUri
}
