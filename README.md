# Movie API

A REST API for managing movies, built with Go and [Gorilla Mux](https://github.com/gorilla/mux). The project includes a Dockerized PostgreSQL development database and supports creating, reading, updating, and deleting movies.

## Features

- List all movies
- Get a movie by ID
- Create a movie
- Replace an existing movie
- Delete a movie
- JSON request and response bodies
- Appropriate HTTP status codes for common outcomes
- PostgreSQL 17 development environment with Docker Compose
- Environment-based application configuration
- Persistent PostgreSQL storage using `pgx`
- Automatic database table creation at application startup

## Requirements

- Go 1.22 or later
- Docker with Docker Compose

## Run locally

Clone the repository and enter the project directory:

```bash
git clone https://github.com/Pratiksable/movie-api.git
cd movie-api
```

Enter the backend module and download its dependencies:

```bash
cd server
go mod download
```

Create `server/.env`:

```env
PORT=8000
DATABASE_URL=postgres://root:root@localhost:5432/movie_db?sslmode=disable
```

Do not add spaces around `=`. The `.env` file is ignored by Git because it can contain local credentials.

Start PostgreSQL:

```bash
docker compose up -d
docker compose ps
```

Load the environment variables and start the API:

```bash
set -a
source .env
set +a
go run ./cmd/api
```

The API will be available at `http://localhost:8000`.

## Project structure

```text
movie-api/
├── server/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go
│   │   ├── database/
│   │   │   └── postgres.go
│   │   ├── handler/
│   │   │   └── movie_handler.go
│   │   ├── model/
│   │   │   └── movie_model.go
│   │   ├── repository/
│   │   │   └── movie_repository.go
│   │   ├── router/
│   │   │   └── movie_router.go
│   │   └── service/
│   │       └── movie_service.go
│   ├── docker-compose.yaml
│   ├── go.mod
│   └── go.sum
├── .gitignore
├── LICENSE
└── README.md
```

- `cmd/api` is the application entry point and starts the HTTP server.
- `internal/config` reads the port and database URL from environment variables.
- `internal/database` creates the connection pool and ensures the database table exists.
- `internal/router` registers the API endpoints and connects them to handlers.
- `internal/handler` processes HTTP requests and responses.
- `internal/model` defines the movie and director data structures.
- `internal/repository` contains the PostgreSQL CRUD queries.
- `internal/service` is reserved for business logic as the application grows.
- `docker-compose.yaml` runs the local PostgreSQL database.

The request flow is:

```text
HTTP request → Router → Handler → Repository → PostgreSQL
```

## PostgreSQL with Docker

The Compose service runs PostgreSQL 17 with these local development settings:

| Setting | Value |
| --- | --- |
| Container | `movie-api-postgres` |
| Host | `localhost` |
| Port | `5432` |
| Database | `movie_db` |
| User | `root` |

View the database logs:

```bash
docker compose logs postgres
```

Stop the containers while preserving database data:

```bash
docker compose down
```

The named Docker volume `movie_postgres_data` preserves PostgreSQL data between container restarts.

When the API starts, it opens a PostgreSQL connection pool and runs `CREATE TABLE IF NOT EXISTS`. No separate migration command is currently required.

## Movie format

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "isbn": "4368827",
  "title": "Movie 1",
  "director": {
    "firstname": "Pratik",
    "lastName": "Sable"
  }
}
```

PostgreSQL generates a UUID when a movie is created. When updating a movie, the ID is taken from the URL.

## API endpoints

| Method | Endpoint | Description | Success status |
| --- | --- | --- | --- |
| `GET` | `/movies` | Get all movies | `200 OK` |
| `GET` | `/movies/{id}` | Get one movie | `200 OK` |
| `POST` | `/movies` | Create a movie | `201 Created` |
| `PUT` | `/movie/{id}` | Replace a movie | `200 OK` |
| `DELETE` | `/movies/{id}` | Delete a movie | `204 No Content` |

Unknown movie IDs return `404 Not Found`. Invalid JSON sent to the create or update endpoints returns `400 Bad Request`.

## Example requests

Get all movies:

```bash
curl http://localhost:8000/movies
```

Get one movie:

```bash
curl http://localhost:8000/movies/550e8400-e29b-41d4-a716-446655440000
```

Create a movie:

```bash
curl -X POST http://localhost:8000/movies \
  -H "Content-Type: application/json" \
  -d '{
    "isbn": "1234567",
    "title": "Avengers",
    "director": {
      "firstname": "Anthony",
      "lastName": "Russo"
    }
  }'
```

Replace a movie:

```bash
curl -X PUT http://localhost:8000/movie/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "isbn": "9876543",
    "title": "Updated Movie",
    "director": {
      "firstname": "Joe",
      "lastName": "Russo"
    }
  }'
```

Because this endpoint uses `PUT`, provide all movie fields in the request body.

Delete a movie:

```bash
curl -X DELETE http://localhost:8000/movies/550e8400-e29b-41d4-a716-446655440000
```

A successful deletion returns an empty response with status `204 No Content`.

## Notes

- Movies are stored in PostgreSQL and remain available after the Go server restarts.
- The movie list starts empty when a new Docker volume is created; add records with `POST /movies`.
- Removing the Docker volume permanently removes the local database data.
- This project is intended for learning Go HTTP APIs and REST fundamentals.

## Development checks

Run these commands from the `server` directory:

```bash
go test ./...
go vet ./...
```
