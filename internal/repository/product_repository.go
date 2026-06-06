package repository

import "game-mini/internal/models"

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindById(id int) (*models.Product, error)
	Create(product models.Product) error
	Update(product models.Product) error
	Delete(id int) error
}