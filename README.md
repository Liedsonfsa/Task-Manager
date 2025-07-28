# 🎯 Task Manager API

This is an API for a Task Manager developed in Go. It provides a RESTful API for task CRUD operations, with JWT authentication and a MySQL database.

## 📍 Endpoints

<!-- ### Authentication

| Method | Endpoint   | Description        |
|--------|------------|--------------------|
| POST   | /register  | Register user      |
| POST   | /login     | User login         |

#### `POST /register`  -->


### Tasks

| Method | Endpoint     | Description                   |
|--------|--------------|-------------------------------|
| GET    | /tasks       | List all tasks                |
| POST   | /tasks       | Create a new task             |
| GET    | /tasks/{id}  | Get task details              |
| PUT    | /tasks/{id}  | Update a task                 |
| DELETE | /tasks/{id}  | Delete a task                 |

#### `GET /tasks`

**Response**

```json
[
    {
        "id": 1,
        "title": "Learn Go",
        "description": "Deepen my knowledge in Golang",
        "status": "Pending",
        "created_at": "2025-04-11T14:16:56-03:00",
        "updated_at": "2025-04-11T14:16:56-03:00"
    },
    {
        "id": 2,
        "title": "Build REST API",
        "description": "Create a RESTful API using Gin framework",
        "status": "In Progress",
        "created_at": "2025-04-15T09:30:00-03:00",
        "updated_at": "2025-04-16T10:45:00-03:00"
    },
    {
        "id": 3,
        "title": "Write Unit Tests",
        "description": "Implement unit tests for the task service",
        "status": "Pending",
        "created_at": "2025-04-17T13:20:00-03:00",
        "updated_at": "2025-04-17T13:20:00-03:00"
    },
    {
        "id": 4,
        "title": "Refactor Code",
        "description": "Improve code readability and structure",
        "status": "Completed",
        "created_at": "2025-04-18T08:00:00-03:00",
        "updated_at": "2025-04-18T11:10:00-03:00"
    }
]
```

#### `POST /tasks`

**Request**

```json
[
    {
        "title": "Set Up MySQL",
        "description": "Configure MySQL database for development",
        "status": "Pending",
    }
]
```

**Response**

```json
[
    {
        "id": 5,
        "title": "Set Up MySQL",
        "description": "Configure MySQL database for development",
        "status": "Pending",
        "created_at": "2025-04-10T15:45:00-03:00",
        "updated_at": "2025-04-10T16:00:00-03:00"
    }
]
```

#### `GET /tasks/:id`

**Response**

```json
[
    {
        "id": 1,
        "title": "Learn Go",
        "description": "Deepen my knowledge in Golang",
        "status": "Pending",
        "created_at": "2025-04-11T14:16:56-03:00",
        "updated_at": "2025-04-11T14:16:56-03:00"
    }
]
```

#### `PUT /tasks/{id}

**Request**

```json
{
    "status": "In Progress"
}
```

**Response**

```json
[
    {
        "id": 5,
        "title": "Set Up MySQL",
        "description": "Configure MySQL database for development",
        "status": "In Progress",
        "created_at": "2025-04-10T15:45:00-03:00",
        "updated_at": "2025-04-11T16:00:00-03:00"
    }
]
```

## 🌐 Technologies Used

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://gin-gonic.com/)
[![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=for-the-badge&logo=mysql&logoColor=white)](https://www.mysql.com/)
[![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white)](https://jwt.io/)