# todo-app
Todo app API built with Go

1. Run with Docker Compose
   
   `docker-compose build`
   
   `docker-compose up`

3. Create Database -> todo_app

4. Seed Database
   
    `go run ./migrate/migrate.go`

6. Restart services
   
   `docker-compose up`

8. Access APIs from localhost:8080
