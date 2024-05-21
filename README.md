# API SPBE

API pemetaan arsitektur SPBE Kota Madiun

API pemetaan data arsitektur SPBE

### Kebutuhan

- Go versi 1.22
- MySQL
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [Make untuk OS Windows](https://gnuwin32.sourceforge.net/packages/make.htm)

### Cara install
<<<<<<< HEAD

- buat database bernama db_spbe
- install golang-migrate (khusus macos / linux)

```sh
make install-migrate
```

- install golang-migrate di cmd (khusus windows)

```sh
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

- migrasi database make os

```sh
make db-migrate
```

- migrasi database windows
- buka terminal ketikkan perintah berikut:

```sh
=======
- buat database bernama db_spbe
- install golang-migrate (khusus macos / linux)

``` sh
make install-migrate
```
- install golang-migrate di cmd (khusus windows)
``` sh
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
- migrasi database make os

``` sh
make db-migrate
```
- migrasi database windows
- buka terminal ketikkan perintah berikut:
``` sh
>>>>>>> 3078a786bb955faf4686788a1bd201aa4bba5757
migrate -path db/migrations -database "mysql://root@tcp(localhost:3306)/db_spbe" up
```

### Run server
<<<<<<< HEAD

- running untuk macos

```sh
make run
```

- running untuk windows
  masuk ke directory cmd

```sh
cd cmd
```

ketikkan perintah:

```sh
=======
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
>>>>>>> 3078a786bb955faf4686788a1bd201aa4bba5757
go run main.go
```

untuk menghentikan server, tekan Ctrl + c
contoh request ada di folder 'example'
<<<<<<< HEAD

- install golang-migrate (khusus windows)
  open cmd dan ketikkan:

```sh
=======
- install golang-migrate (khusus windows)
open cmd dan ketikkan:
``` sh
>>>>>>> 3078a786bb955faf4686788a1bd201aa4bba5757
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Target perilisan

### versi 0 (1 hari)

- [ ] pembuatan dokumentasi
- [ ] API sederhana menampilkan respon ("SPBE KOTA MADIUN")
- [ ] unit tes
- [ ] menjalankan di server

### versi 1 (3 hari)

- [ ] ....
