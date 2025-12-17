# Movie API

A RESTful API for managing users and their movie collections, built with Go and Gin framework.

## Features

- User management (create, list users)
- Movie management (add, delete movies for users)
- Retrieve movies associated with a specific user
- MySQL database integration
- Environment-based configuration

## Tech Stack

- **Language**: Go 1.25.5
- **Framework**: Gin
- **Database**: MySQL
- **Environment Management**: godotenv

## Prerequisites

- Go 1.25.5 or later
- MySQL 8.0 or later
- Git

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/masaya-nishimura-09/movie-api.git
   cd movie-api
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Set up the database:

   - Create a MySQL database named `movie_api`
   - Run the SQL script in `db` to create tables:

4. Create a `.env` file in the root directory with your database credentials:
   ```
   DBUSER=your_mysql_username
   DBPASS=your_mysql_password
   ```

## Usage

1. Start the server:

   ```bash
   go run cmd/main.go
   ```

2. The API will be available at `http://localhost:8080`

## API Endpoints

### Users

- **GET /user** - Retrieve all users
- **POST /user** - Create a new user
  - Body: `{"name": "string"}`
- **GET /user/:id** - Retrieve movies for a specific user

### Movies

- **POST /movie** - Add a movie to a user's collection
  - Body: `{"userId": number, "imdbId": "string", "title": "string", "director": "string", "year": number, "rating": number, "comment": "string"}`
- **DELETE /movie/:id** - Delete a movie by ID

## Data Models

### User

```json
{
  "id": 1,
  "name": "John Doe",
  "updatedAt": "2025-12-17T10:00:00Z",
  "createdAt": "2025-12-17T10:00:00Z"
}
```

### Movie

```json
{
  "id": 1,
  "userId": 1,
  "imdbId": "tt0054130",
  "title": "La Notte",
  "director": "Michelangelo Antonioni",
  "year": 1961,
  "rating": 4,
  "comment": "A masterpiece",
  "updatedAt": "2025-12-17T10:00:00Z",
  "createdAt": "2025-12-17T10:00:00Z"
}
```

## Database Schema

The application uses the following tables:

- `users`: Stores user information
- `movies`: Stores movie information linked to users

## Error Handling

The API returns JSON error responses with the following structure:

```json
{
  "error": "ERROR_CODE",
  "message": "Human-readable error message"
}
```

Common error codes:

- `INVALID_USER_ID`: User ID must be a number
- `INVALID_MOVIE_ID`: Movie ID must be a number
- `INVALID_JSON`: Request body must be valid JSON
- `INTERNAL_SERVER_ERROR`: Server-side error occurred

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests (if any)
5. Submit a pull request

## License

This project is licensed under the terms specified in the LICENSE file.
