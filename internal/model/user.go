package model

type User struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    UpdatedAt string `json:"updatedAt"`
    CreatedAt string `json:"createdAt"`
}

type Movie struct {
    Id int64 `json:"id"`
    UserId int `json:"userId"`
    ImdbId string `json:"imdbId"`
    Title string `json:"title"`
    Director string `json:"director"`
    Year int `json:"year"`
    Rating int `json:"rating"`
    Comment string `json:"comment"`
    UpdatedAt string `json:"updatedAt"`
    CreatedAt string `json:"createdAt"`
}

type UserMovies struct {
    User User
    Movies []Movie
}
