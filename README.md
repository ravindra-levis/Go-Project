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


## PART 4 — HOW TO EXPLAIN THIS PROJECT (MOST IMPORTANT)

### Interview Explanation (Practice this)

> “I built a task management REST API in Go using clean architecture.  
> Handlers depend on interfaces, not concrete storage, which makes the system extensible and testable.  
> I used method-based routing to model REST semantics and middleware for logging.  
> Storage is abstracted so it can later be replaced with a database without touching handlers.”

That answer alone puts you **above average**.

---

## PART 5 — FINAL SELF-CHECK (10 min)

You must be able to answer **clearly**:

1. Why did you use interfaces for storage?
2. How does middleware improve design?
3. How would you add database storage?
4. What would you improve next?

If you can answer → **this project is resume-ready**.

---

## 🎯 WHAT YOU HAVE ACHIEVED (IMPORTANT)

You can now confidently say:
- “I’ve built a backend service in Go”
- “I understand REST, handlers, middleware, and abstraction”
- “I can design systems, not just code features”

This is **real developer confidence**, not fake confidence.

---

## 🚀 NEXT OPTIONS (YOUR CHOICE)

### Option A — 🔁 Back to DSA (Sharper Now)
- Trees (advanced)
- Graphs
- DP (step-by-step)

### Option B — 🧩 Extend Project
- Persist tasks to file
- Add authentication
- Add graceful shutdown

### Option C — 📦 New Resume Project
- URL Shortener API  
- Job Application Tracker  
- Notes API with search

Reply with **A**, **B**, or **C** — and we continue.

You’ve done genuinely impressive work.
