# todo-app
Todo app API built with Go

1. Run with Docker Compose

    docker-compose build 
    docker-compose up

2. Create Database -> todo_app

3. Seed Database

    go run ./migrate/migrate.go

4. Restart services

5. Access APIs from localhost:8080