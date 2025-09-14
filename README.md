
```markdown
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
- [ ] Integrasi dengan Frontend Vue.js


## 🛠️ Tech Stack
- **Backend**: Golang (Gin Framework)
- **Database**: MySQL / PostgreSQL (via GORM)
- **Auth**: JWT (JSON Web Token)
- **Payment**: Midtrans / Xendit (planned)
- **Frontend**: Vue.js (planned)


## 📂 Struktur Project
```

Project-golang/
├── handler/         # Handler (controller) untuk API
├── helper/          # Utility & helper function
├── campaign/        # Modul campaign
├── user/            # Modul user & auth
├── transaction/     # Modul transaksi & pembayaran
├── auth/            # JWT auth middleware
├── main.go          # Entry point aplikasi





## 🚀 Cara Menjalankan
### 1. Clone Repo
```bash
git clone https://github.com/endrycofr/Project-golang.git
cd Project-golang
````

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

* [x] Integrasi Payment Gateway
* [x] Dashboard untuk Admin
* [ ] Vue.js Frontend
* [ ] Dockerization
* [ ] Deployment ke Cloud (AWS / GCP)



