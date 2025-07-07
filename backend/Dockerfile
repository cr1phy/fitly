FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o fitly ./cmd/fitly
EXPOSE 8080
CMD ["./fitly"]