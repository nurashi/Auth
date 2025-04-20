FROM golang:1.24-alpine

RUN apk add --no-cache git build-base

WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080


# real time update, with using compileDaemon
CMD CompileDaemon --build="go build -buildvcs=false -o main ." --command="./main" --exclude-dir=".git,tmp" --graceful-kill