FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod tidy

# Compilar el proyecto
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

ENV MONGO_USERNAME=admin
ENV MONGO_PASSWORD=password
ENV MONGO_HOST=mongodb
ENV MONGO_PORT=27017
ENV MONGO_DATABASE=test
ENV PORT=8080

EXPOSE 8080

CMD ["./main"]