FROM golang:1.25.2

WORKDIR /app

COPY . .

RUN go mod tidy

CMD [ "go", "run", "main.go" ]