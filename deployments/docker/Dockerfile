# Stage 1: Build
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache  ca-certificates tzdata

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o app .

# Stage 2: Run
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /build/app /app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app"]