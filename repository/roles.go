package repository

import (
	"github.com/GilangAndhika/bukuin_be/models"
	"gorm.io/gorm"
)

func GetAllRoles(db *gorm.DB) ([]models.Roles, error) {
	var roles []models.Roles
	// Mengambil semua data role dari database
	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func GetRoleByID(db *gorm.DB, id string) (models.Roles, error) {
	var role models.Roles
	// Mengambil data role berdasarkan ID
	if err := db.First(&role, id).Error; err != nil {
		return role, err
	}
	return role, nil
}

func CreateRole(db *gorm.DB, role models.Roles) error {
	// Menambahkan data role ke database
	if err := db.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRole(db *gorm.DB, id string, UpdatedRoles models.Roles) error {
	// Memperbarui data role di database
	if err := db.Model(&models.Roles{}).Where("id_role = ?", id).Updates(UpdatedRoles).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRole(db *gorm.DB, id string) error {
	// Menghapus data role dari database
	if err := db.Delete(&models.Roles{}, id).Error; err != nil {
		return err
	}
	return nil
}
