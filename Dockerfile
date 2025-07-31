FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build -o auth-service main.go

# Minimal runtime image
FROM gcr.io/distroless/static

# 👇 Copy CA certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/auth-service /

CMD ["/auth-service"]
