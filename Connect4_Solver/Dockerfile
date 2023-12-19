FROM golang:1.21.5

COPY . /app

WORKDIR /app

CMD cd Magic && make && cd .. && go run Solver.go