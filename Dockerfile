# Build stage
FROM golang:1.24 as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o auth-service .

# Run stage
FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/auth-service /app/auth-service
ENTRYPOINT ["/app/auth-service"]

