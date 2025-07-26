# Task Manager API

This is an API for a Task Manager developed in Go. It provides a RESTful API for task CRUD operations, with JWT authentication and a MySQL database.

## Endpoints

### Authentication

| Method | Endpoint   | Description        |
|--------|------------|--------------------|
| POST   | /register  | Register user      |
| POST   | /login     | User login         |

### Tasks

| Method | Endpoint     | Description                   |
|--------|--------------|-------------------------------|
| GET    | /tasks       | List all tasks                |
| POST   | /tasks       | Create a new task             |
| GET    | /tasks/{id}  | Get task details              |
| PUT    | /tasks/{id}  | Update a task                 |
| DELETE | /tasks/{id}  | Delete a task                 |

## Technologies Used

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://gin-gonic.com/)
[![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=for-the-badge&logo=mysql&logoColor=white)](https://www.mysql.com/)
[![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white)](https://jwt.io/)