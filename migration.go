package database

import (
    "myapp/models"
)

func Migrate() {
    DB.AutoMigrate(&models.User{}, &models.Photo{})
}