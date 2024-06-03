package controllers

import (
    "myapp/database"
    "myapp/helpers"
    "myapp/models"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/asaskevich/govalidator"
)

func Register(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if !govalidator.IsEmail(input.Email) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
        return
    }

    hashedPassword, _ := helpers.HashPassword(input.Password)
    input.Password = hashedPassword

    if err := database.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": input})
}

func Login(c *gin.Context) {
    var input models.User
    var storedUser models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if !govalidator.IsEmail(input.Email) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
        return
    }

    if err := database.DB.Where("email = ?", input.Email).First(&storedUser).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
        return
    }

    if !helpers.CheckPasswordHash(input.Password, storedUser.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := helpers.GenerateJWT(storedUser.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func UpdateUser(c *gin.Context) {
    var user models.User
    if err := database.DB.Where("id = ?", c.Param("userId")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
        return
    }

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Save(&user)
    c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
    var user models.User
    if err := database.DB.Where("id = ?", c.Param("userId")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
        return
    }

    database.DB.Delete(&user)
    c.JSON(http.StatusOK, gin.H{"data": true})
}
