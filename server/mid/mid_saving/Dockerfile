FROM golang:1.21.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./mid-saving

EXPOSE 50050
EXPOSE 50051
EXPOSE 50052

CMD ["./mid-saving"]

