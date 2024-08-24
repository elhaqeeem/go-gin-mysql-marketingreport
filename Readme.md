<h2>MARKETING REPORT</h2>
GOLANG (GIN) MICROSERVICE - MYSQL
<br>
## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Struktur files](#Struktur files)
<br>#Struktur files

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
