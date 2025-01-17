FROM golang:1.22.0-alpine3.19

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main cmd/main.go

EXPOSE 5640
CMD ["./main"]