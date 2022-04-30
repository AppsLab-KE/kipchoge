package storage

import (
	"appslab-ke/kipchoge-go/internal/models"
	"appslab-ke/kipchoge-go/internal/platform"
	"errors"
	"os"
)

type Storage interface {
	ConnectDb() error
	FindAllJokeCategories() ([]models.JokeCategory, error)
	Ping() error
	Defer() error
}

func NewStorage() (Storage, error) {
	selectedDb := os.Getenv("SELECTED_DB")
	switch selectedDb {
	case "mongodb":
		storage, err := platform.NewMongoDbConn()
		if err != nil {
			return nil, err
		}
		return storage, nil
	}

	return nil, errors.New("not database set")
}
