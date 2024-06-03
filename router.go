package router

import (
    "myapp/controllers"
    "myapp/middlewares"
    "github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
    user := r.Group("/users")
    {
        user.POST("/register", controllers.Register)
        user.GET("/login", controllers.Login)
        user.PUT("/:userId", middlewares.Auth(), controllers.UpdateUser)
        user.DELETE("/:userId", middlewares.Auth(), controllers.DeleteUser)
    }

    photo := r.Group("/photos")
    {
        photo.POST("/", middlewares.Auth(), controllers.CreatePhoto)
        photo.GET("/", middlewares.Auth(), controllers.GetPhotos)
        photo.PUT("/:photoId", middlewares.Auth(), controllers.UpdatePhoto)
        photo.DELETE("/:photoId", middlewares.Auth(), controllers.DeletePhoto)
    }
}