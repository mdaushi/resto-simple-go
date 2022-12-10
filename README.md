# RESTO SIMPLE GO

## TEKNOLOGI YANG DIGUNAKAN

- [gin](https://gin-gonic.com/) salah satu framework go dengan pengguna banyak dan dukungan komunitas yang baik
- [gorm](https://gorm.io/) ORM lib untuk golang
- [viper](https://github.com/spf13/viper) mengefisiensikan config go

## Struktur folder

- cmd - Aplikasi utama untuk proyek ini.
  - main.go
- pkg - modular service yang dibutuhkan
  - common
    - config - config project
    - db - conifg db
    - envs - env
    - models - type model db
    - responses - helper untuk responses api
    - seeds - seeder model
  - orders - partial module orders
  - ..... - partial module lainnya

## Endpoint

ambil semua user

    GET - http://localhost:3000/users


ambil semua menu

    GET - http://localhost:3000/menus

tampilkan stok

    GET - http://localhost:3000/menus/stock

buat orderan

    POST - http://localhost:3000/orders

lihat semua orderan

    GET - http://localhost:3000/orders

print struk

    GET - http://localhost:3000/receipts/<id orderan>/print

laporan bulanan

    GET - http://localhost:3000/reports/monthly

laporan tahunan

    GET - http://localhost:3000/reports/yearly

## Cara jalankan
1. setup dan jalankan postgres
2. paste config string Postgres ke .env di pkg/common/env
3. go run cmd/main.go

*pada saat pertama kali dijalankan seeder user dan menu langsung runnig. jika tidak ingin jalankan seeder komentar kodenya di pkg/common/db/db.go