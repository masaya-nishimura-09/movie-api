package handler

import (
	"net/http"
    "log"
    "strconv"

    "github.com/gin-gonic/gin"
	"github.com/masaya-nishimura-09/movie-api/internal/repository"
	"github.com/masaya-nishimura-09/movie-api/internal/model"
)

func UserMovies(c *gin.Context) {
    userId := c.Param("id")

	userIdInt, err := strconv.Atoi(userId)
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "INVALID_USER_ID",
            "message": "userId must be a number",
        })
        return
    }

    movies, err := repository.UserMovies(userIdInt)
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "INTERNAL_SERVER_ERROR",
            "message": "failed to fetch movies",
        })
        return
    }
    c.JSON(http.StatusOK, movies)
}

func AddMovie(c *gin.Context) {
    var movie model.Movie
    if err := c.ShouldBindJSON(&movie); err != nil {
        log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "INVALID_JSON",
            "message": "request body must be valid JSON",
        })
        return
    }

    id, err := repository.AddMovie(&movie)
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "INTERNAL_SERVER_ERROR",
            "message": "failed to add a movie",
        })
        return
    }
    movie.Id = id
    c.JSON(http.StatusOK, movie)
}

func DeleteMovie(c *gin.Context) {
    id := c.Param("id")

	idInt, err := strconv.Atoi(id)
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "INVALID_MOVIE_ID",
            "message": "movieId must be a number",
        })
        return
    }

    dbErr := repository.DeleteMovie(idInt)
    if dbErr != nil {
        log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "INTERNAL_SERVER_ERROR",
            "message": "failed to delete a movie",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{ "movieId": idInt })
}
