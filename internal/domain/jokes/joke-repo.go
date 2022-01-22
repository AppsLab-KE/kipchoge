package jokes

import (
	"appslab-ke/kipchoge-go/internal/models"
	"appslab-ke/kipchoge-go/internal/storage"
)

type JokeRepo struct {
	client storage.Storage
}

func NewJokeRepo(db storage.Storage) JokeRepo {
	return JokeRepo{
		client: db,
	}
}

func (r JokeRepo) GetJokeCategories() ([]models.JokeCategory, error) {
	return r.client.FindAllJokeCategories()
}