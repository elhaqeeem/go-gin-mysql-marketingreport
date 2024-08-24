<h2>MARKETING REPORT</h2>
GOLANG (GIN) MICROSERVICE - MYSQL
<br>

## Struktur File

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



## Prerequisites

Before running this project, ensure you have the following installed:

- [Go](https://golang.org/dl/)
- [Node.js](https://nodejs.org/) (includes npm)

## Backend Setup

1. Navigate to the backend directory:

   ```bash
   cd go-backend
