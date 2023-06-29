FROM golang:alpine

WORKDIR /app
RUN mkdir "/build"

ADD . .

RUN go mod download

RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build -o /build/app" -command="/build/app"