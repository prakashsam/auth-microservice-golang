FROM golang:1.24.5 as builder

WORKDIR /app
COPY . .

ENV CGO_ENABLED=0
RUN go build -o auth-service main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/auth-service .

CMD ["/app/auth-service"]
