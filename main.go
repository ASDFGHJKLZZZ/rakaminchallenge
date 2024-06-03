package main

import (
    "myapp/database"
    "myapp/router"
    "github.com/gin-gonic/gin"
)

func main() {
    database.InitDB()
    r := gin.Default()
    router.SetupRouter(r)
    r.Run(":8080")
}