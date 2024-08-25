


  
# Project Title
<h2>MARKETING REPORT</h2>
## Description
GOLANG (GIN) MICROSERVICE - MYSQL
## Table of Contents
- [Installation](#installation)
- [Usage](#usage)


## Installation
To install and set up the project, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/elhaqeeem/go-resto-mysql.git

## Usage

2. Copy env file 
   ```sh
    cp .env .env.backup

3. Running local --> delete comment in 
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


4. Deploy to aws or etc --> upload or bulk environment in setting deployment

5. Command to running 
   ```sh
   go mod tidy
   go run main.go




