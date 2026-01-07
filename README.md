# Project Crowdfunding CMS - Golang Backend

🚀 Backend API untuk platform **Crowdfunding CMS** yang dibangun dengan **Golang (Gin Framework)**.  
Project ini nantinya akan terintegrasi dengan frontend menggunakan **Vue.js** untuk memberikan pengalaman pengguna yang lebih interaktif.

## 📌 Fitur (Progress)

- [x] API Authentication (Login & Register)
- [x] CRUD User (Create, Read, Update, Delete)
- [x] CRUD Campaign (Create, Update, Get Campaigns)
- [x] Payment Gateway Integration (Midtrans / Xendit)
- [x] Admin Dashboard API
- [ ] Unit Testing
- [x] Integrasi dengan Frontend Vue.js

## 🛠️ Tech Stack

- **Backend**: Golang (Gin Framework)
- **Database**: MySQL / PostgreSQL (via GORM)
- **Auth**: JWT (JSON Web Token)
- **Payment**: Midtrans / Xendit (planned)
- **Frontend**: Vue.js (planned)

## 📂 Struktur Project
```bash
Project-golang/
├── handler/ # Handler (controller) untuk API
├── helper/ # Utility & helper function
├── campaign/ # Modul campaign
├── user/ # Modul user & auth
├── transaction/ # Modul transaksi & pembayaran
├── auth/ # JWT auth middleware
├── main.go # Entry point aplikasi
```
## 🚀 Cara Menjalankan

### 1. Clone Repo

```bash
git clone https://github.com/endrycofr/Project-golang.git
cd Project-golang
```

### 2. Setup Database

Buat database di MySQL/PostgreSQL sesuai konfigurasi di file `main.go`.

### 3. Jalankan Aplikasi

```bash
go run main.go
```

API akan berjalan di `http://localhost:8080`.

## 📡 Endpoint API (sementara)

| Method | Endpoint            | Deskripsi              |
| ------ | ------------------- | ---------------------- |
| POST   | `/api/v1/users`     | Register user baru     |
| POST   | `/api/v1/sessions`  | Login user (JWT token) |
| GET    | `/api/v1/campaigns` | List campaign          |
| POST   | `/api/v1/campaigns` | Tambah campaign baru   |

## 📅 Roadmap

- [x] Integrasi Payment Gateway
- [x] Dashboard untuk Admin
- [x] Vue.js Frontend
- [ ] Dockerization
- [ ] Deployment ke Cloud (AWS / GCP)

---

## 🧩 Tech Stack

### Frontend

Frontend aplikasi ini dibangun menggunakan Nuxt 4 dengan pendekatan SSR-first, modern, dan scalable. Arsitektur dirancang agar mudah dikembangkan untuk kebutuhan funding / CMS serta terintegrasi dengan backend Golang (Gin) menggunakan JWT.

- **Nuxt 4**
- TypeScript
- Tailwind CSS
- `useFetch` (built-in Nuxt)
- Pinia (state management)
- Server-Side Rendering (SSR) enabled
- Auth menggunakan Pinia + Nuxt Middleware

### Backend

- Golang
- Gin Framework
- JWT Authentication
- MySQL
- REST API

---

## 🔑 Konsep Web Ini

### 1. **Nuxt 4 default pakai SSR**

- Request API dilakukan di **server Nuxt**

### 2. Menggunakan `useFetch`, **bukan axios**

- Lebih optimal untuk SSR
- Auto hydration
- Best practice Nuxt 4

## 🚦 Next Step

1. Detail campaign (`/campaigns/:slug`)
2. Login JWT (POST `/sessions`)
3. Simpan token di state / cookie
4. Attach `Authorization: Bearer <token>`
5. Middleware auth Nuxt

## ✨ Tujuan

Integrasi **Nuxt 4 + Golang (Gin)**:

- Clean
- SSR-friendly
- Type-safe
- Scalable untuk aplikasi **funding / CMS**
