package repository

import (
    "database/sql"
    "os"

    "github.com/go-sql-driver/mysql"
    "github.com/masaya-nishimura-09/movie-api/internal/model"
)

func database() (*sql.DB, error) {
    var db *sql.DB

    // Capture connection properties.
    cfg := mysql.NewConfig()
    cfg.User = os.Getenv("DBUSER")
    cfg.Passwd = os.Getenv("DBPASS")
    cfg.Net = "tcp"
    cfg.Addr = "localhost:3306"
    cfg.DBName = "movie_api"

    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        return nil, err
    }

    pingErr := db.Ping()
    if pingErr != nil {
        return nil, err
    }

    return db, nil
}

func Users() ([]model.User, error) {
    db, err := database()
    if err != nil {
        return nil, err
    }

    users := []model.User{}

    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var user model.User
        if err := rows.Scan(
            &user.Id, 
            &user.Name, 
            &user.UpdatedAt, 
            &user.CreatedAt,
        ); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func AddUser(user *model.User) (int64, error) {
    db, err := database()
    if err != nil {
        return 0, err
    }

    result, err := db.Exec("INSERT INTO users (name) VALUES (?)", user.Name)
    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func AddMovie(movie *model.Movie) (int64, error) {
    db, err := database()
    if err != nil {
        return 0, err
    }

    result, err := db.Exec("INSERT INTO movies (user_id, imdb_id, title, director, year, rating, comment) VALUES (?, ?, ?, ?, ?, ?, ?)", movie.UserId, movie.ImdbId, movie.Title, movie.Director, movie.Year, movie.Rating, movie.Comment)
    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func UserMovies(id int) ([]model.Movie, error) {
    db, err := database()
    if err != nil {
        return nil, err
    }

    movies := []model.Movie{}

    rows, err := db.Query("SELECT * FROM movies WHERE user_id = ?", id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var movie model.Movie
        if err := rows.Scan(
            &movie.Id, 
            &movie.UserId, 
            &movie.ImdbId, 
            &movie.Title, 
            &movie.Director, 
            &movie.Year, 
            &movie.Rating, 
            &movie.Comment, 
            &movie.UpdatedAt, 
            &movie.CreatedAt,
        ); err != nil {
            return nil, err
        }
        movies = append(movies, movie)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return movies, nil
}

func DeleteMovie(id int) (error) {
    db, err := database()
    if err != nil {
        return err
    }

    _, deleteErr := db.Exec("DELETE FROM movies WHERE id = ?", id)
    if deleteErr != nil {
        return err
    }

    return nil
}
