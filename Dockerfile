FROM golang:1.24.0

RUN apt-get update && apt-get upgrade -y \

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

CMD ["go", "run", "main.go"]
