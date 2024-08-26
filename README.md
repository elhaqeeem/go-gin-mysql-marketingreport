# MARKETING REPORT

## Deskripsi
Proyek ini adalah aplikasi mikroservices yang dibangun dengan Golang (Gin) , menggunakan MySQL sebagai basis datanya dan React sebagai frondend 

## Daftar Pustaka
- [Instalasi](#instalasi)
- [Struktur File](#struktur-file)
- [Penggunaan](#penggunaan)


<img src="media/1.png" alt="Alt Text" width="500" height="300">


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
   

   Rute API

   Marketing

   Buat Marketing
   Endpoint: POST /marketing
   Handler: handlers.CreateMarketing(db)
   Deskripsi: Menambahkan catatan marketing baru.
      ```json
      {
            "name": "Alfandy"
      }
      ```
   Dapatkan Marketing berdasarkan ID
   Endpoint: GET /marketing/:id
   Handler: handlers.GetMarketing(db)
   Deskripsi: Mengambil data marketing berdasarkan ID.
      ```json
      {
            "id": 1,
            "name": "Alfandy"
      }
      ```
   Dapatkan Semua Marketing
   Endpoint: GET /marketing
   Handler: handlers.GetAllMarketing(db)
   Deskripsi: Mengambil semua data marketing.
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
   Perbarui Marketing
   Endpoint: PUT /marketing/:id
   Handler: handlers.UpdateMarketing(db)
   Deskripsi: Memperbarui data marketing berdasarkan ID.
      ```json
      {
               "name": "Alfandy"
      }
      ```
   Hapus Marketing
   Endpoint: DELETE /marketing/:id
   Handler: handlers.DeleteMarketing(db)
   Deskripsi: Menghapus data marketing berdasarkan ID.
      ```json
      {
               "name": "Alfandy"
      }
      ```


   Penjualan

   Buat Penjualan
   Endpoint: POST /penjualan
   Handler: handlers.CreatePenjualan(db)
   Deskripsi: Menambahkan catatan penjualan baru.
     
      ```json
      {
            "MarketingID": 1,
            "Date": "2023-05-22",
            "CargoFee": 25000,
            "TotalBalance": 3000000,
      }
      ```

   Dapatkan Penjualan berdasarkan ID
   Endpoint: GET /penjualan/:id
   Handler: handlers.GetPenjualan(db)
   Deskripsi: Mengambil data penjualan berdasarkan ID.
      ```json
      {
            "id": 1,
            "TransactionNumber": "TRX001",
            "MarketingID": 1,
            "Date": "2023-05-22T00:00:00Z",
            "CargoFee": 25000,
            "TotalBalance": 3000000,
            "GrandTotal": 3025000
      }
      ```
   Dapatkan Semua Penjualan
   Endpoint: GET /penjualan
   Handler: handlers.GetallPenjualan(db)
   Deskripsi: Mengambil semua data penjualan.
      ```json
      [
         {
            "id": 1,
            "TransactionNumber": "TRX001",
            "MarketingID": 1,
            "Date": "2023-05-22T00:00:00Z",
            "CargoFee": 25000,
            "TotalBalance": 3000000,
            "GrandTotal": 3025000
         },
         {
            "id": 2,
            "TransactionNumber": "TRX002",
            "MarketingID": 3,
            "Date": "2023-05-22T00:00:00Z",
            "CargoFee": 25000,
            "TotalBalance": 320000,
            "GrandTotal": 345000
         },
         {
            "id": 3,
            "TransactionNumber": "TRX003",
            "MarketingID": 1,
            "Date": "2023-05-22T00:00:00Z",
            "CargoFee": 0,
            "TotalBalance": 65000000,
            "GrandTotal": 65000000
         },
         {
            "id": 4,
            "TransactionNumber": "TRX004",
            "MarketingID": 1,
            "Date": "2023-05-23T00:00:00Z",
            "CargoFee": 10000,
            "TotalBalance": 70000000,
            "GrandTotal": 70010000
         },
         {
            "id": 5,
            "TransactionNumber": "TRX005",
            "MarketingID": 2,
            "Date": "2023-05-23T00:00:00Z",
            "CargoFee": 10000,
            "TotalBalance": 80000000,
            "GrandTotal": 80010000
         },
         {
            "id": 6,
            "TransactionNumber": "TRX006",
            "MarketingID": 3,
            "Date": "2023-05-23T00:00:00Z",
            "CargoFee": 12000,
            "TotalBalance": 44000000,
            "GrandTotal": 44012000
         },
         {
            "id": 7,
            "TransactionNumber": "TRX007",
            "MarketingID": 1,
            "Date": "2023-06-01T00:00:00Z",
            "CargoFee": 0,
            "TotalBalance": 75000000,
            "GrandTotal": 75000000
         },
         {
            "id": 8,
            "TransactionNumber": "TRX008",
            "MarketingID": 2,
            "Date": "2023-06-02T00:00:00Z",
            "CargoFee": 0,
            "TotalBalance": 85000000,
            "GrandTotal": 85000000
         },
         {
            "id": 9,
            "TransactionNumber": "TRX009",
            "MarketingID": 2,
            "Date": "2023-06-01T00:00:00Z",
            "CargoFee": 0,
            "TotalBalance": 175000000,
            "GrandTotal": 175000000
         },
         {
            "id": 10,
            "TransactionNumber": "TRX010",
            "MarketingID": 3,
            "Date": "2023-06-01T00:00:00Z",
            "CargoFee": 0,
            "TotalBalance": 75000000,
            "GrandTotal": 75000000
         },
         {
            "id": 11,
            "TransactionNumber": "TRX011",
            "MarketingID": 2,
            "Date": "2023-06-01T00:00:00Z",
            "CargoFee": 0,
            "TotalBalance": 750020000,
            "GrandTotal": 750020000
         },
         {
            "id": 12,
            "TransactionNumber": "TRX012",
            "MarketingID": 3,
            "Date": "2023-06-01T00:00:00Z",
            "CargoFee": 0,
            "TotalBalance": 130000000,
            "GrandTotal": 130000000
         }
      ]
      ```
   Perbarui Penjualan
   Endpoint: PUT /penjualan/:id
   Handler: handlers.UpdatePenjualan(db)
   Deskripsi: Memperbarui data penjualan berdasarkan ID.
      ```json
      {
         {
            "id": 1,
            "MarketingID": 1,
            "Date": "2023-05-22",
            "CargoFee": 25000,
            "TotalBalance": 3000000,
         }
      }
      ```
   Hapus Penjualan
   Endpoint: DELETE /penjualan/:id
   Handler: handlers.DeletePenjualan(db)
   Deskripsi: Menghapus data penjualan berdasarkan ID.
      ```json
      {
         {
            "id": 1,
            "MarketingID": 1,
            "Date": "2023-05-22",
            "CargoFee": 25000,
            "TotalBalance": 3000000,
         }
      }
      ```
   Pembayaran dan Komisi
      
   Dapatkan Komisi
   Endpoint: GET /komisi
   Handler: handlers.GetKomisi(db)
   Deskripsi: Mengambil data komisi.
      ```json
      [
         {
            "marketing_id": 1,
            "bulan": "2023-05",
            "omzet": 138035000,
            "komisi_persen": 2.5,
            "komisi_nominal": 3450875
         },
         {
            "marketing_id": 3,
            "bulan": "2023-05",
            "omzet": 44357000,
            "komisi_persen": 0,
            "komisi_nominal": 0
         },
         {
            "marketing_id": 2,
            "bulan": "2023-05",
            "omzet": 80010000,
            "komisi_persen": 0,
            "komisi_nominal": 0
         },
         {
            "marketing_id": 1,
            "bulan": "2023-06",
            "omzet": 75000000,
            "komisi_persen": 0,
            "komisi_nominal": 0
         },
         {
            "marketing_id": 2,
            "bulan": "2023-06",
            "omzet": 1010020000,
            "komisi_persen": 10,
            "komisi_nominal": 101002000
         },
         {
            "marketing_id": 3,
            "bulan": "2023-06",
            "omzet": 205000000,
            "komisi_persen": 5,
            "komisi_nominal": 10250000
         }
      ]
      ```
   Buat Pembayaran
   Endpoint: POST /pembayaran
   Handler: handlers.CreatePembayaran(db)
   Deskripsi: Membuat catatan pembayaran baru.
      ```json
      {
            "marketing_id": 2,
            "amount": 650000,
            "jumlah_angsuran": "3",
            "payment_method": "credit"
      }
      ```
   Dapatkan Semua Pembayaran
   Endpoint: GET /pembayaran
   Handler: handlers.GetPembayaran(db)
   Deskripsi: Mengambil semua data pembayaran.
      ```json
      [
  
         {
            "id": 10,
            "marketing_id": 2,
            "amount": 250000,
            "payment_date": "2023-05-22T00:00:00Z",
            "status": "completed",
            "payment_method": "debet"
         },
         {
            "id": 11,
            "marketing_id": 2,
            "amount": 650000,
            "payment_date": "2023-05-22T00:00:00Z",
            "status": "pending",
            "payment_method": "credit"
         },
         {
            "id": 12,
            "marketing_id": 2,
            "amount": 30000,
            "payment_date": "2024-08-26T00:00:00Z",
            "status": "completed",
            "payment_method": "debet"
         },
         {
            "id": 13,
            "marketing_id": 2,
            "amount": 30000,
            "payment_date": "2024-08-26T00:00:00Z",
            "status": "completed",
            "payment_method": "debet"
         },
         {
            "id": 14,
            "marketing_id": 3,
            "amount": 130000,
            "payment_date": "2024-08-26T00:00:00Z",
            "status": "pending",
            "payment_method": "credit"
         },
         {
            "id": 15,
            "marketing_id": 3,
            "amount": 300000,
            "payment_date": "2024-08-26T00:00:00Z",
            "status": "pending",
            "payment_method": "credit"
         },
         {
            "id": 16,
            "marketing_id": 3,
            "amount": 300000,
            "payment_date": "2024-08-26T00:00:00Z",
            "status": "pending",
            "payment_method": "credit"
         },
         {
            "id": 17,
            "marketing_id": 1,
            "amount": 100000,
            "payment_date": "2024-08-26T00:00:00Z",
            "status": "pending",
            "payment_method": "credit"
         },
         {
            "id": 18,
            "marketing_id": 1,
            "amount": 500000,
            "payment_date": "2024-08-26T00:00:00Z",
            "status": "pending",
            "payment_method": "credit"
         }
      ]
      ```
   Dapatkan Semua Angsuran untuk Pembayaran Tertentu
   Endpoint: GET /angsuran/:pembayaran_id
   Handler: handlers.GetAllAngsuran(db)
   Deskripsi: Mengambil semua data angsuran untuk pembayaran tertentu berdasarkan ID pembayaran.
      ```json
      {
            "id": 18,
            "marketing_id": 1,
            "amount": 500000,
            "payment_date": "2024-08-26T00:00:00Z",
            "status": "pending",
            "payment_method": "credit"
      }
      ```
   Cek Status Angsuran Pertama untuk Pembayaran Tertentu
   Endpoint: GET /angsuran/status/:pembayaran_id
   Handler: handlers.CheckInstallmentStatus(db)
   Deskripsi: Memeriksa status pembayaran dari angsuran pertama untuk ID pembayaran tertentu.
     ```json
      {
            "angsuran_ke": 1,
            "jumlah_angsuran": 100000,
            "tanggal_pembayaran": "2024-08-26",
            "status": "telah dibayar"
      }
     ```

