# MARKETING REPORT

## Deskripsi
Proyek ini adalah aplikasi mikroservices yang dibangun dengan Golang (Gin) dan menggunakan MySQL sebagai basis datanya.

## Daftar Pustaka
- [Instalasi](#instalasi)
- [Penggunaan](#penggunaan)
- [Struktur File](#struktur-file)
- [Endpoint](#endpoint)

## Instalasi
Untuk menginstal dan mengatur proyek, ikuti langkah-langkah berikut:

1. **Clone repositori:**
   ```sh
   git clone https://github.com/elhaqeeem/go-resto-mysql.git

2. **salin file .env :**
   ```sh
   cp .env.backup .env

3. **install migrate go untuk windows / mac0s / linux::**
   ```sh
   scoop install migrate / 
   brew install golang-migrate / 
   $ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
   $ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
   $ apt-get update
   $ apt-get install -y migrate


3. **Perintah untuk menjalankan::**
   ```sh
   migrate -path db/migrations -database "mysql://username:password@tcp(host:port)/namadb" up
   go run main.go
