FROM golang:1.21

WORKDIR /app

COPY go.mod .
COPY ./server .

RUN go get -u github.com/lib/pq

RUN go get -u github.com/gin-gonic/gin

RUN go build -o main



CMD ["./main"]