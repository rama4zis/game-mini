package repository

import "game-mini/internal/models"

type GameRepository interface {
	FindAll() ([]models.Game, error)
	FindById(id int) (*models.Game, error)
	Create(game models.Game) error
	Update(game models.Game) error
	Delete(id int) error
}

