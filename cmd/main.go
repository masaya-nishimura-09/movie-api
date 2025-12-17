package main

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
	"github.com/masaya-nishimura-09/movie-api/internal/handler"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println(err)
    }

    router := gin.Default()
    router.GET("/user", handler.Users)
    router.GET("/user/:id", handler.UserMovies)
    router.POST("/user", handler.AddUser)
    router.POST("/movie", handler.AddMovie)
    router.DELETE("/movie/:id", handler.DeleteMovie)
    router.Run("localhost:8080")
}


