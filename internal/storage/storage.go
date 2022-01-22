package storage

import "appslab-ke/kipchoge-go/internal/models"

type Storage interface {
	ConnectDb() error
	FindAllJokeCategories() ([]models.JokeCategory, error)
}