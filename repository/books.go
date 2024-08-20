package repository

import (
	"github.com/gryzlegrizz/bukuin_be/models"
	"gorm.io/gorm"
)

func GetAllBooks(db *gorm.DB) ([]models.Book, error) {
	var books []models.Book
	// Mengambil semua data buku dari database
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}