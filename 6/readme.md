# Setup Instructions
Untuk menjalankan API secara lokal, mesin lokal harus memiliki Docker Engine yang sudah terinstall.

Setelah itu, buatlah sebuah `.env` file sesuai dengan key yang ada pada file `.env.example`.

Setelah itu, silakan masuk ke directory `6` dan jalankan command ini pada terminal.
```
docker-compose up --build -d
```

Setelah menjalankan command ini, API akan berjalan di mesin lokal.

 # API Testing dan Cara Penggunaan
Dokumentasi dari URL API yang ada dan example request-response dapat dilihat pada tautan berikut: https://documenter.getpostman.com/view/12531688/2s935kPkY8

Workspace untuk Postman API dapat diakses pada tautan berikut: https://www.postman.com/richard-here/workspace/shopping-cart-api/collection/12531688-ffc54ac4-06ba-415b-9d75-492c3dd47f41?action=share&creator=12531688.

**Catatan penting**: workspace pada Postman API berisi pre-request scripts dan test scripts yang mengkonfigurasi nilai pada body request yang akan digunakan untuk melakukan automated API endpoint testing. Oleh karena itu, agar behavior expected, **endpoints tidak boleh dijalankan secara manual**. Untuk menjalankan automated API testing, pastikan untuk melakukan hal berikut.
- Pastikan webserver app sudah berjalan di mesin lokal menggunakan Docker
- Set environment menjadi `Shopping Cart API Test` pada workspace
- Klik kanan pada collection `Shopping Cart API Test`
- Klik pilihan `Run Collection`
- Pastikan nilai `Iterations` adalah `1` dan `Delay` adalah `50`
- Klik tombol `Run Shopping Cart API Test`
- Lihat hasil testing
- Apabila testing ingin diulang, jangan lupa untuk menghapus dulu test data yang sudah dimasukkan di DB dengan menjalankan `DELETE FROM "products"` pada terminal Docker container PostgreSQL yang berjalan. Tujuannya adalah agar data yang existing tidak mengganggu expected behavior dari test suites.

# Unit Testing
File unit testing terdapat pada folder `cart-api/test`. Untuk menjalankan unit testing dari root directory `6`, jalankan perintah ini pada terminal.
```
go test -v ./cart-api/test
```
Hasil testing dapat dilihat pada terminal.