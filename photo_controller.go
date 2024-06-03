package controllers

import (
    "myapp/database"
    "myapp/models"
    "net/http"
    "github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
    var photo models.Photo
    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Create(&photo)
    c.JSON(http.StatusOK, gin.H{"data": photo})
}

func GetPhotos(c *gin.Context) {
    var photos []models.Photo
    database.DB.Find(&photos)
    c.JSON(http.StatusOK, gin.H{"data": photos})
}

func UpdatePhoto(c *gin.Context) {
    var photo models.Photo
    if err := database.DB.Where("id = ?", c.Param("photoId")).First(&photo).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Photo not found"})
        return
    }

    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Save(&photo)
    c.JSON(http.StatusOK, gin.H{"data": photo})
}

func DeletePhoto(c *gin.Context) {
    var photo models.Photo
    if err := database.DB.Where("id = ?", c.Param("photoId")).First(&photo).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Photo not found"})
        return
    }

    database.DB.Delete(&photo)
    c.JSON(http.StatusOK, gin.H{"data": true})
}
