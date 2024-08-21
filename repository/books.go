package repository

import (
	"github.com/gryzlegrizz/bukuin_be/models"
	"gorm.io/gorm"
)

func GetAllBooks(db *gorm.DB) ([]models.Books, error) {
	var books []models.Books
	// Mengambil semua data buku dari database
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookByID(db *gorm.DB, id string) (models.Books, error) {
	var books models.Books
	// Mengambil data buku berdasarkan ID
	if err := db.First(&books, id).Error; err != nil {
		return books, err
	}
	return books, nil
}

func CreateBook(db *gorm.DB, book *models.Books) error {
	// Menambahkan data buku ke database
	if err := db.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, id string, UpdatedBooks models.Books) error {
	// Memperbarui data buku di database
	if err := db.Model(&models.Books{}).Where("id_book = ?", id).Updates(UpdatedBooks).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBook(db *gorm.DB, id string) error {
	// Menghapus data buku dari database
	if err := db.Delete(&models.Books{}, id).Error; err != nil {
		return err
	}
	return nil
}