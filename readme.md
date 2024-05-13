# Panduan Penggunaan Aplikasi Enigma Laundry

Aplikasi Enigma Laundry merupakan aplikasi pengelolaan data pelanggan dan transaksi laundry Anda. Aplikasi ini berbasis konsol dan terhubung dengan PostgreSQL. Berikut adalah panduan penggunaannya:

1. Mengelola Data Pelanggan

1.1. Melihat Data Pelanggan:

- Jalankan aplikasi pada terminal "go run main.go".
- Pilih "Data Pelanggan" (1).
- Pilih "Lihat Data Pelanggan" (1).
- Data pelanggan (ID, nama, nomor HP) akan ditampilkan.

1.2. Menambahkan Pelanggan Baru:

- Jalankan aplikasi pada terminal "go run main.go".
- Pilih "Data Pelanggan" (1).
- Pilih "Tambah Data Pelanggan Baru" (2).
- Masukkan ID, nama, dan nomor HP pelanggan baru.
- Konfirmasi penambahan data (Y/N).

1.3. Mengubah Data Pelanggan:

- Jalankan aplikasi pada terminal "go run main.go".
- Pilih "Data Pelanggan" (1).
- Pilih "Perbarui Data Pelanggan" (3).
- Masukkan ID pelanggan yang ingin diubah.
- Masukkan data baru (nama, nomor HP) atau tekan Enter jika tidak ingin mengubahnya.
- Konfirmasi perubahan data (Y/N).

1.4. Menghapus Data Pelanggan:

- Jalankan aplikasi pada terminal "go run main.go".
- Pilih "Data Pelanggan" (1).
- Pilih "Hapus Data Pelanggan" (4).
- Masukkan ID pelanggan yang ingin dihapus.
- Konfirmasi penghapusan data (Y/N).

2. Mengelola Data Transaksi

2.1.Melihat Transaksi:

- Jalankan aplikasi pada terminal "go run main.go".
- Pilih "Data Transaksi" (2).
- Pilih "Lihat Data Transaksi" (1).
- Pilih cara pencarian (ID pelanggan, nama pelanggan, atau ID transaksi).
- Masukkan informasi pencarian.
- Data transaksi (nama pelanggan, layanan, kuantitas, pegawai, tanggal masuk, total harga) akan ditampilkan.

2.2. Menambahkan Transaksi Baru:

- Jalankan aplikasi pada terminal "go run main.go".
- Pilih "Data Transaksi" (2).
- Pilih "Tambah Transaksi" (2).
- Masukkan ID pelanggan dan pegawai.
- Masukkan tanggal perkiraan selesai (DD-MM-YYYY).
- Pilih layanan dan masukkan kuantitas.
- Masukkan status pembayaran (Lunas/Belum Lunas).
- Konfirmasi penambahan transaksi (Y/N).


Catatan:
- Pastikan Anda berada di direktori aplikasi saat menjalankan perintah "go run main.go"
- Jika data yang dimasukkan tidak valid atau tidak ditemukan, aplikasi akan memberikan keterangan input yang benar.