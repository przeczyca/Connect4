FROM golang:1.21.5

COPY . /app

WORKDIR /app

EXPOSE 8080

CMD cd internal/Magic && make && cd ../.. &&  go run cmd/Connect4_Solver/main.go