# =========================
# Stage 1: Builder
# =========================
FROM golang:1.24-alpine AS builder

WORKDIR /build

# Install CA certificates untuk HTTPS request
RUN apk add --no-cache ca-certificates tzdata

# Download dependency
COPY go.mod go.sum ./
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off
RUN go mod download

# Copy source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o app .

# =========================
# Stage 2: Runtime
# =========================
FROM scratch

# Copy CA certificates (penting untuk HTTPS, Minio, API, dll)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy timezone
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy binary
COPY --from=builder /build/app /app

# Expose port
EXPOSE 8080

# Non root user
USER 1001

ENTRYPOINT ["/app"]