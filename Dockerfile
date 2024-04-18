FROM golang:1.22-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
VOLUME /root/internal/database/
COPY --from=builder /app/main .
COPY --from=builder /app/internal/frontend/templates/ ./internal/frontend/templates
COPY --from=builder /app/internal/frontend/static/ ./internal/frontend/static
CMD ["./main"]