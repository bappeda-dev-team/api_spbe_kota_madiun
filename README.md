# API SPBE
API pemetaan arsitektur SPBE Kota Madiun

API pemetaan data arsitektur SPBE

### Kebutuhan
- Go versi 1.22
- MySQL
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [Make untuk OS Windows](https://gnuwin32.sourceforge.net/packages/make.htm)

### Cara install
- buat database bernama db_spbe
- install golang-migrate (khusus macos / linux)

``` sh
make install-migrate
```

- migrasi database

``` sh
make db-migrate
```
- install golang-migrate (khusus windows)
open cmd dan ketikkan:
``` sh
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Run server
- running untuk macos
``` sh
make run
```
- running untuk windows
masuk ke directory cmd
``` sh
cd cmd
```
ketikkan perintah:
``` sh
go run main.go
```

untuk menghentikan server, tekan Ctrl + c
contoh request ada di folder 'example'

