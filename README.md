# SimpleTodoAPI

SimpleTodoAPI is a RESTful web service built using the Gin framework in Go. It provides endpoints to manage a simple Todo application, including creating, reading, updating, and deleting tasks.

## Features

- Create new tasks
- Retrieve all tasks or a specific task
- Update existing tasks
- Delete tasks
- Database integration for persistent storage

## Project Structure

```
.
├── api/
│   ├── handlers.go       # Contains the API handlers
│   ├── model.go          # Defines the data models
├── cmd/
│   ├── main.go           # Entry point of the application
├── db/
│   ├── config.go         # Database configuration
│   ├── schema.sql        # SQL schema for database setup
├── routes/
│   ├── routes.go         # API route definitions
├── tests/                # Contains test files
├── go.mod                # Go module file
├── go.sum                # Go dependencies
├── LICENSE               # License file
├── README.md             # Project documentation
```

## Prerequisites

- Go 1.20 or later
- A running instance of a database (e.g., PostgreSQL, MySQL, SQLite)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/MukalDadhwal/SimpleTodoAPI.git
   cd SimpleTodoAPI
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up the database:

   - Create a database and configure the connection in `db/config.go`.
   - Run the SQL schema in `db/schema.sql` to set up the tables.

## Running the Application

To start the server, run:

```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080` by default.

## API Endpoints

### Base URL

`http://localhost:8080/api/v1`

### Endpoints

| Method | Endpoint       | Description              |
|--------|----------------|--------------------------|
| GET    | `/todos`       | Retrieve all tasks       |
| GET    | `/todos/:id`   | Retrieve a specific task |
| POST   | `/todos`       | Create a new task        |
| PUT    | `/todos/:id`   | Update an existing task  |
| DELETE | `/todos/:id`   | Delete a task            |

### Example Request

#### Create a New Task

```bash
curl -X POST http://localhost:8080/api/v1/todos \
-H "Content-Type: application/json" \
-d '{
  "title": "Buy groceries",
  "description": "Milk, Bread, Eggs"
}'
```

## Testing

Run the tests using:

```bash
go test ./tests/...
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
