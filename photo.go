package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type Photo struct {
    ID        uint      `gorm:"primary_key" json:"id"`
    Title     string    `json:"title"`
    Caption   string    `json:"caption"`
    PhotoUrl  string    `json:"photo_url"`
    UserID    uint      `json:"user_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}