FROM golang:1.21.5

WORKDIR /app

COPY go.mod go.sum ./
COPY wait-kafka.sh ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./consumer_system

RUN chmod +x ./wait-kafka.sh

ENTRYPOINT ["./wait-kafka.sh"]