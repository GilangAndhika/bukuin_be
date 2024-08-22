package repository

import (
    "github.com/gryzlegrizz/bukuin_be/models"
    "gorm.io/gorm"
)

func UsernameExists(db *gorm.DB, username string) (bool, error) {
    var user models.Users
    result := db.Where("username = ?", username).First(&user)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return false, nil
        }
        return false, result.Error
    }
    return true, nil
}
