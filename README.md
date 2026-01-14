# Task Management REST API

A simple REST API built in Go for managing tasks.

## Features

- Create, read, update, and delete tasks
- Clean architecture with handlers, models, and storage layers
- In-memory storage with interface-based design
- JSON-based REST API

## Installation

```bash
go run main.go
```

Server runs on `http://localhost:8080`

## API Endpoints

### Create Task
```bash
POST /tasks
Content-Type: application/json

{
  "title": "Learn Go",
  "done": false
}
```

### List All Tasks
```bash
GET /tasks
```

### Update Task
```bash
PUT /tasks/{id}
Content-Type: application/json

{
  "done": true
}
```

### Delete Task
```bash
DELETE /tasks/{id}
```

## Project Structure

```
task-api/
├── main.go              # Entry point
├── handlers/
│   └── task.go         # HTTP handlers
├── models/
│   └── task.go         # Task model
└── storage/
    ├── store.go        # Storage interface
    └── memory.go       # In-memory implementation
```

## Example Usage

```bash
# Create a task
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Complete project", "done": false}'

# List all tasks
curl http://localhost:8080/tasks

# Update task
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"done": true}'

# Delete task
curl -X DELETE http://localhost:8080/tasks/1
```
