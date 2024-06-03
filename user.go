package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type User struct {
    ID        uint      `gorm:"primary_key" json:"id"`
    Username  string    `gorm:"not null" json:"username"`
    Email     string    `gorm:"unique;not null" json:"email"`
    Password  string    `gorm:"not null" json:"password"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Photos    []Photo   `gorm:"constraint:OnDelete:CASCADE" json:"photos"`
}