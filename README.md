# Task Management REST API (Go)

A production-style REST API built in Go using the standard library.

## Features
- Create, list, update, and delete tasks
- Clean architecture (handlers, models, storage)
- In-memory storage with interface-based abstraction
- Middleware-based request logging
- JSON-based REST API

## Endpoints

### Create Task
POST /tasks
```json
{ "title": "Learn Go", "done": false }
List Tasks
GET /tasks
Update Task
PUT /tasks/{id}
{ "done": true }
Delete Task
DELETE /tasks/{id}

Design Decisions
- Used interfaces to decouple storage from handlers
- Method-based routing for REST compliance
- Middleware for cross-cutting concerns


## HOW TO EXPLAIN THIS PROJECT (MOST IMPORTANT)

### Interview Explanation (Practice this)

> “I built a task management REST API in Go using clean architecture.  
> Handlers depend on interfaces, not concrete storage, which makes the system extensible and testable.  
> I used method-based routing to model REST semantics and middleware for logging.  
> Storage is abstracted so it can later be replaced with a database without touching handlers.”


---

You must be able to answer:

1. Why did you use interfaces for storage?
2. How does middleware improve design?
3. How would you add database storage?
4. What would you improve next?
