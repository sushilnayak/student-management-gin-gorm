FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o student-service

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/student-service .
COPY config.yaml .

EXPOSE 8080

CMD ["./student-service"]