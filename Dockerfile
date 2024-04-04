# Build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM scratch
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

