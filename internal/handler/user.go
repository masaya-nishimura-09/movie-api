package handler

import (
	"net/http"
    "log"

    "github.com/gin-gonic/gin"
	"github.com/masaya-nishimura-09/movie-api/internal/repository"
	"github.com/masaya-nishimura-09/movie-api/internal/model"
)

func Users(c *gin.Context) {
    users, err := repository.Users()
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "INTERNAL_SERVER_ERROR",
            "message": "failed to fetch users",
        })
        return
    }

    c.JSON(http.StatusOK, users)
}

func AddUser(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "INVALID_JSON",
            "message": "request body must be valid JSON",
        })
        return
    }

    id, err := repository.AddUser(&user)
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "INTERNAL_SERVER_ERROR",
            "message": "failed to create user",
        })
        return
    }
    user.Id = id
    c.JSON(http.StatusOK, user)
}
