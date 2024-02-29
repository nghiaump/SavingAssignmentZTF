FROM golang:1.21.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY wait-es.sh ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./core-saving

# Cấp quyền thực thi cho script wait-es.sh
RUN chmod +x ./wait-es.sh

# Thiết lập ENTRYPOINT để chạy script kiểm tra sẵn sàng của Elasticsearch trước khi chạy ứng dụng
ENTRYPOINT ["./wait-es.sh"]

EXPOSE 50052

# Lenh nay da duoc chen luon vao ENTRYPOINT
# CMD ["./core-saving"]
