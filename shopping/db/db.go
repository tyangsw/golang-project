package db

import (
	"staging/shopping/models"
)

func LoadItem(id int) *models.Item {
	return &models.Item{
		Price: 3.42,
	}
}
