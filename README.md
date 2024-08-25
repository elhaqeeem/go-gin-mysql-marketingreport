# MARKETING REPORT

## Deskripsi
   Golang (gin) microservices + mysql

## Daftar Pustaka
- [Install](#Install)
- [Penggunaan](#Penggunaan)
- [Struktur-file](#Struktur file)
- [Endpoint](#Endpoint)

## Install
To install and set up the project, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/elhaqeeem/go-resto-mysql.git

## Penggunaan

2. Copy env file 
   ```sh
    cp .env.backup .env

3. Running local --> delete comment in 
   ```sh
  


4. Deploy to aws or etc --> upload or bulk environment in setting deployment

5. Command to running 
   ```sh
   go mod tidy
   go run main.go

## Struktur file
1. Command to running 
   ```sh
   markerting-report/
   
    ├── main.go
    ├── config/
    │   └── config.go
    ├── db/
    │   └── migration/
    │       └── migration.sql
    ├── models/
    │   ├── marketing.go
    │   ├── penjualan.go
    │   └── pembayaran.go
    │   └── komisi.go
    ├── handlers/
    │   ├── komisi.go
    │   └── pembayaran.go
    └── utils/
        └── response.go

## Endpoint
   




