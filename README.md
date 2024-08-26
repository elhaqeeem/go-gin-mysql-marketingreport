# MARKETING REPORT

## Deskripsi
Proyek ini adalah aplikasi mikroservices yang dibangun dengan Golang (Gin) , menggunakan MySQL sebagai basis datanya dan React sebagai frondend 

## Daftar Pustaka
- [Instalasi](#instalasi)
- [Struktur File](#struktur-file)
- [Penggunaan](#penggunaan)

## Instalasi
   Untuk menginstal dan mengatur proyek, ikuti langkah-langkah berikut:

1. **Clone repositori:**
   ```sh
   git clone https://github.com/elhaqeeem/go-gin-mysql-marketingreport.git
2. **salin file .env :**
   ```sh
   cp .env.backup .env

3. **install migrate go untuk windows / mac0s / linux::**
   ```sh
   $ scoop install migrate (untuk windows)

   $ brew install golang-migrate (untuk macos)
   
   $ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
   $ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
   $ apt-get update
   $ apt-get install -y migrate
   (untuk linux)

4. **Perintah untuk menjalankan::**
   ```sh
   migrate -path db/migrations -database "mysql://username:password@tcp(host:port)/namadb" up
   go run main.go

## Struktur File
1. **Marketing report:**
   ```sh

   ├── main.go
   ├── config/
   │   └── config.go
   ├── db/
   │   └── migration/
   │       └── migration.sql
   ├── models/
   │   ├── marketing.go
   │   ├── penjualan.go
   │   ├── pembayaran.go
   │   └── komisi.go
   ├── handlers/
   │   ├── marketing.go
   │   ├── penjualan.go
   │   ├── pembayaran.go
   │   └── komisi.go
   └── utils/
      └── response.go

  


## Penggunaan
1. **Route Api:**
   ```sh

   Rute API

   Marketing

   Buat Marketing
   Endpoint: POST /marketing
   Handler: handlers.CreateMarketing(db)
   Deskripsi: Menambahkan catatan marketing baru.

   Dapatkan Marketing berdasarkan ID
   Endpoint: GET /marketing/:id
   Handler: handlers.GetMarketing(db)

   Deskripsi: Mengambil data marketing berdasarkan ID.
   Dapatkan Semua Marketing
   Endpoint: GET /marketing
   ```json
      [
         {
            "id": 1,
            "name": "Alfandy"
         },
         {
            "id": 2,
            "name": "Mery"
         },
         {
            "id": 3,
            "name": "Danang"
         }
         ]

      ```
         
   Handler: handlers.GetAllMarketing(db)
   Deskripsi: Mengambil semua data marketing.

   Perbarui Marketing
   Endpoint: PUT /marketing/:id
   Handler: handlers.UpdateMarketing(db)
   Deskripsi: Memperbarui data marketing berdasarkan ID.
   
   Hapus Marketing
   Endpoint: DELETE /marketing/:id
   Handler: handlers.DeleteMarketing(db)
   Deskripsi: Menghapus data marketing berdasarkan ID.

   Penjualan

   Buat Penjualan
   Endpoint: POST /penjualan
   Handler: handlers.CreatePenjualan(db)
   Deskripsi: Menambahkan catatan penjualan baru.

   Dapatkan Penjualan berdasarkan ID
   Endpoint: GET /penjualan/:id
   Handler: handlers.GetPenjualan(db)
   Deskripsi: Mengambil data penjualan berdasarkan ID.

   Dapatkan Semua Penjualan
   Endpoint: GET /penjualan
   Handler: handlers.GetallPenjualan(db)
   Deskripsi: Mengambil semua data penjualan.

   Perbarui Penjualan
   Endpoint: PUT /penjualan/:id
   Handler: handlers.UpdatePenjualan(db)
   Deskripsi: Memperbarui data penjualan berdasarkan ID.

   Hapus Penjualan
   Endpoint: DELETE /penjualan/:id
   Handler: handlers.DeletePenjualan(db)
   Deskripsi: Menghapus data penjualan berdasarkan ID.

   Pembayaran dan Komisi

   Dapatkan Komisi
   Endpoint: GET /komisi
   Handler: handlers.GetKomisi(db)
   Deskripsi: Mengambil data komisi.

   Buat Pembayaran
   Endpoint: POST /pembayaran
   Handler: handlers.CreatePembayaran(db)
   Deskripsi: Membuat catatan pembayaran baru.

   Dapatkan Semua Pembayaran
   Endpoint: GET /pembayaran
   Handler: handlers.GetPembayaran(db)
   Deskripsi: Mengambil semua data pembayaran.

   Dapatkan Semua Angsuran untuk Pembayaran Tertentu
   Endpoint: GET /angsuran/:pembayaran_id
   Handler: handlers.GetAllAngsuran(db)
   Deskripsi: Mengambil semua data angsuran untuk pembayaran tertentu berdasarkan ID pembayaran.

   Cek Status Angsuran Pertama untuk Pembayaran Tertentu
   Endpoint: GET /angsuran/status/:pembayaran_id
   Handler: handlers.CheckInstallmentStatus(db)
   Deskripsi: Memeriksa status pembayaran dari angsuran pertama untuk ID pembayaran tertentu.