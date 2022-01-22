package services

import "appslab-ke/kipchoge-go/internal/models"

type JokeService interface {
	GetJokeCategories() ([]models.JokeCategory, error)
}

