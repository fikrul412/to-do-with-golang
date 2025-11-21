# Todo API Documentation

This is the RESTful API for the Todo application built with Go (Gin) and PostgreSQL.

---

## Base URL

http://localhost:3000/api

pgsql
Copy code

---

## Todos Endpoints

### 1. List Todos

**GET /todos**

Retrieve all todos with optional pagination, search, and category filter.

**Query Parameters:**

| Parameter    | Type   | Description                                      |
|--------------|--------|--------------------------------------------------|
| page         | int    | Page number (default: 1)                        |
| limit        | int    | Items per page (default: 10)                    |
| search       | string | Filter by title or description containing text  |
| category_id  | int    | Filter by specific category ID                  |

**Response:**

```json
{
  "data": [
    {
      "id": 1,
      "title": "Test Todo",
      "description": "desc",
      "completed": false,
      "category_id": 1,
      "priority": "high",
      "due_date": "2025-11-21T10:00:00Z"
    }
  ]
}
```

### 2. Get Todo by ID

**GET /todos/:id**

Retrieve a specific todo by its ID.

**Response:**

```json
{
  "data": {
    "id": 1,
    "title": "Test Todo",
    "description": "desc",
    "completed": false,
    "category_id": 1,
    "priority": "high",
    "due_date": "2025-11-21T10:00:00Z"
  }
}
```
### 3. Create Todo

**POST /todos**

Create a new todo.

**Request Body:**

```json
{
  "title": "Test Todo",
  "description": "desc",
  "priority": "high",
  "category_id": 1,
  "due_date": "2025-11-21T10:00:00Z"
}
```
**Response:**

```json
{
  "data": {
    "id": 1,
    "title": "Test Todo",
    "description": "desc",
    "completed": false,
    "category_id": 1,
    "priority": "high",
    "due_date": "2025-11-21T10:00:00Z"
  }
}

```

### 4. Update Todo

**PUT /todos/:id**

Update an existing todo.

**Request Body:**

```json
{
  "title": "Updated Todo",
  "description": "updated description",
  "priority": "medium",
  "category_id": 1,
  "due_date": "2025-11-22T12:00:00Z"
}
```
**Response:**

```json
{
  "data": {
    "id": 1,
    "title": "Updated Todo",
    "description": "updated description",
    "completed": false,
    "category_id": 1,
    "priority": "medium",
    "due_date": "2025-11-22T12:00:00Z"
  }
}
```

### 5. Delete Todo

**DELETE /todos/:id**

Delete a todo by ID.


**Response:**

```json
{
  "message": "Deleted"
}
```

### 6. Toggle Complete

**PATCH /todos/:id/complete**

Toggle the completion status of a todo.

**Response:**

```json
{
  "data": {
    "id": 1,
    "title": "Updated Todo",
    "description": "updated description",
    "completed": true,
    "category_id": 1,
    "priority": "medium",
    "due_date": "2025-11-22T12:00:00Z"
  }
}

```

## Categories Endpoints

### 1. List Categories

**GET /categories**

Retrieve all categories.


**Response:**

```json
{
  "data": [
    {
      "id": 1,
      "name": "Work",
      "color": "blue"
    }
  ]
}

```

### 2. Create Category

**POST /categories**

Retrieve all categories.

**Request Body:**

```json
{
  "name": "Work",
  "color": "blue"
}

```

**Response:**

```json
{
  "data": {
    "id": 1,
    "name": "Work",
    "color": "blue"
  }
}

```

Notes:
------
* All endpoints are prefixed with /api
* Dates should use ISO 8601 format
* category_id must reference an existing category
* Pagination defaults: page=1, limit=10