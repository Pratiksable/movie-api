# Movie API

A simple REST API for managing movies, built with Go and [Gorilla Mux](https://github.com/gorilla/mux). Movies are stored in memory and the API supports creating, reading, updating, and deleting them.

## Features

- List all movies
- Get a movie by ID
- Create a movie
- Replace an existing movie
- Delete a movie
- JSON request and response bodies
- Appropriate HTTP status codes for common outcomes

## Requirements

- Go 1.22 or later

## Run locally

Clone the repository and enter the project directory:

```bash
git clone https://github.com/Pratiksable/movie-api.git
cd movie-api
```

Enter the backend module, download its dependencies, and start the API:

```bash
cd server
go mod download
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
│   │   ├── handler/
│   │   │   └── movie_handler.go
│   │   ├── model/
│   │   │   └── movie_model.go
│   │   ├── router/
│   │   │   └── movie_router.go
│   │   └── service/
│   │       └── movie_service.go
│   ├── go.mod
│   └── go.sum
├── .gitignore
├── LICENSE
└── README.md
```

- `cmd/api` is the application entry point and starts the HTTP server.
- `internal/router` registers the API endpoints and connects them to handlers.
- `internal/handler` processes HTTP requests and responses.
- `internal/model` defines the movie and director data structures.
- `internal/service` is reserved for business logic as the application grows.

## Movie format

```json
{
  "id": "1",
  "isbn": "4368827",
  "title": "Movie 1",
  "director": {
    "firstname": "Pratik",
    "lastName": "Sable"
  }
}
```

The server generates the `id` when a movie is created. When updating a movie, the ID is taken from the URL.

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
curl http://localhost:8000/movies/1
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
curl -X PUT http://localhost:8000/movie/1 \
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
curl -X DELETE http://localhost:8000/movies/2
```

A successful deletion returns an empty response with status `204 No Content`.

## Notes

- Data is stored in memory and resets whenever the server restarts.
- The in-memory movie list starts empty; create records with `POST /movies`.
- This project is intended for learning Go HTTP APIs and REST fundamentals.

## Development checks

Run these commands from the `server` directory:

```bash
go test ./...
go vet ./...
```
