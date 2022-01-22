package jokes

import "appslab-ke/kipchoge-go/internal/models"

type IJokeRepo interface {
	GetJokeCategories() ([]models.JokeCategory, error)
}