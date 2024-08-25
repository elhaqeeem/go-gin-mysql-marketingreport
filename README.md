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
Struktur File

Struktur direktori untuk proyek ini adalah sebagai berikut:

go
Copy code
marketing-report/
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
│   ├── komisi.go
│   └── pembayaran.go
└── utils/
    └── response.go
Endpoint

1. Dapatkan Komisi
Endpoint: GET /komisi
Deskripsi: Mengambil data komisi.
Respons:
200 OK: Mengembalikan array JSON dari catatan komisi.
2. Buat Pembayaran
Endpoint: POST /pembayaran
Deskripsi: Membuat catatan pembayaran baru.
Body Permintaan:
json
Copy code
{
  "marketing_id": 2,
  "amount": 25000,
  "payment_date": "2023-05-22T00:00:00Z",
  "status": "completed",
  "payment_method": "debit"
}
Respons:
200 OK: Mengembalikan pesan sukses.
400 Bad Request: Mengembalikan pesan kesalahan jika data input tidak valid.
500 Internal Server Error: Mengembalikan pesan kesalahan jika ada masalah di server.
3. Dapatkan Pembayaran
Endpoint: GET /pembayaran
Deskripsi: Mengambil semua catatan pembayaran.
Respons:
200 OK: Mengembalikan array JSON dari catatan pembayaran.
4. Dapatkan Semua Angsuran
Endpoint: GET /angsuran/:pembayaran_id
Deskripsi: Mengambil semua angsuran untuk ID pembayaran tertentu.
Parameter:
pembayaran_id (path): ID dari pembayaran.
Respons:
200 OK: Mengembalikan array JSON dari catatan angsuran untuk pembayaran yang ditentukan.
404 Not Found: Mengembalikan pesan kesalahan jika tidak ada angsuran ditemukan untuk ID pembayaran tersebut.
5. Cek Status Angsuran
Endpoint: GET /angsuran/status/:pembayaran_id
Deskripsi: Memeriksa status pembayaran dari angsuran pertama untuk ID pembayaran tertentu.
Parameter:
pembayaran_id (path): ID dari pembayaran.
Respons:
200 OK: Mengembalikan status dari angsuran pertama.
404 Not Found: Mengembalikan pesan kesalahan jika angsuran tidak ditemukan.
400 Bad Request: Mengembalikan pesan kesalahan jika ID pembayaran tidak valid.
Rute CRUD Marketing
6. Buat Marketing
Endpoint: POST /marketing
Deskripsi: Membuat catatan marketing baru.
Body Permintaan:
json
Copy code
{
  "name": "Nama Marketing",
  "details": "Detail Marketing"
}
Respons:
200 OK: Mengembalikan pesan sukses.
400 Bad Request: Mengembalikan pesan kesalahan jika data input tidak valid.
7. Dapatkan Marketing
Endpoint: GET /marketing/:id
Deskripsi: Mengambil catatan marketing berdasarkan ID.
Parameter:
id (path): ID dari catatan marketing.
Respons:
200 OK: Mengembalikan catatan marketing.
404 Not Found: Mengembalikan pesan kesalahan jika ID marketing tidak ada.
8. Dapatkan Semua Marketing
Endpoint: GET /marketing
Deskripsi: Mengambil semua catatan marketing.
Respons:
200 OK: Mengembalikan array JSON dari catatan marketing.
9. Perbarui Marketing
Endpoint: PUT /marketing/:id
Deskripsi: Memperbarui catatan marketing yang ada.
Parameter:
id (path): ID dari catatan marketing yang akan diperbarui.
Body Permintaan:
json
Copy code
{
  "name": "Nama Marketing Yang Diperbarui",
  "details": "Detail Marketing Yang Diperbarui"
}
Respons:
200 OK: Mengembalikan pesan sukses.
400 Bad Request: Mengembalikan pesan kesalahan jika data input tidak valid.
404 Not Found: Mengembalikan pesan kesalahan jika ID marketing tidak ada.
10. Hapus Marketing
Endpoint: DELETE /marketing/:id
Deskripsi: Menghapus catatan marketing berdasarkan ID.
Parameter:
id (path): ID dari catatan marketing yang akan dihapus.
Respons:
200 OK: Mengembalikan pesan sukses.
404 Not Found: Mengembalikan pesan kesalahan jika ID marketing tidak ada.
Rute CRUD Penjualan
11. Buat Penjualan
Endpoint: POST /penjualan
Deskripsi: Membuat catatan penjualan baru.
Body Permintaan:
json
Copy code
{
  "product_id": 1,
  "quantity": 5,
  "total_price": 50000
}
Respons:
200 OK: Mengembalikan pesan sukses.
400 Bad Request: Mengembalikan pesan kesalahan jika data input tidak valid.
12. Dapatkan Penjualan
Endpoint: GET /penjualan/:id
Deskripsi: Mengambil catatan penjualan berdasarkan ID.
Parameter:
id (path): ID dari catatan penjualan.
Respons:
200 OK: Mengembalikan catatan penjualan.
404 Not Found: Mengembalikan pesan kesalahan jika ID penjualan tidak ada.
13. Dapatkan Semua Penjualan
Endpoint: GET /penjualan
Deskripsi: Mengambil semua catatan penjualan.
Respons:
200 OK: Mengembalikan array JSON dari catatan penjualan.
14. Perbarui Penjualan
Endpoint: PUT /penjualan/:id
Deskripsi: Memperbarui catatan penjualan yang ada.
Parameter:
id (path): ID dari catatan penjualan yang akan diperbarui.
Body Permintaan:
json
Copy code
{
  "product_id": 2,
  "quantity": 3,
  "total_price": 30000
}
Respons:
200 OK: Mengembalikan pesan sukses.
400 Bad Request: Mengembalikan pesan kesalahan jika data input tidak valid.
404 Not Found: Mengembalikan pesan kesalahan jika ID penjualan tidak ada.
15. Hapus Penjualan
Endpoint: DELETE /penjualan/:id
Deskripsi: Menghapus catatan penjualan berdasarkan ID.
Parameter:
id (path): ID dari catatan penjualan yang akan dihapus.
Respons:
200 OK: Mengembalikan pesan sukses.
404 Not Found: Mengembalikan pesan kesalahan jika ID penjualan tidak ada.
