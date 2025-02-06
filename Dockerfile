# Gunakan base image Go
FROM golang:1.20

# Set working directory dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Unduh semua dependency Go
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Expose port yang digunakan aplikasi
EXPOSE 8080

# Jalankan aplikasi saat container dimulai
CMD ["./main"]
