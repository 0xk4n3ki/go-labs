FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o go_k8s

FROM scratch
WORKDIR /app
COPY --from=builder /app/go_k8s .
CMD ["./go_k8s"]