package configs

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"os"
)

const selectedDb = "mongo"

func GetDbConfig() string {
	var dbUri string
	switch selectedDb {
	case "mongo":
		{
			mongoUri, ok := os.LookupEnv("MONGO_URI")
			if !ok {
				panic(errors.New("mongo uri not defined in env"))
			}
			dbUri = mongoUri
		}
	default:
		dbUri = ""
	}

	return dbUri
}

func GetRedisConfig() *redis.Options {
	redisAddr, ok := os.LookupEnv("REDIS_ADDRESS")
	if !ok {
		panic(errors.New("redis address not defined in env"))
		return nil
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

	return &redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	}
}
