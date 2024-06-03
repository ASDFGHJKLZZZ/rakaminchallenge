package database

import (
    "myapp/models"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "log"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open("mysql", "root:@tcp(localhost:3306)/myapp?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        log.Fatal(err)
    }
    DB.AutoMigrate(&models.User{}, &models.Photo{})
}