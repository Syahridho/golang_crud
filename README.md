# Aplikasi Web CRUD Go

Aplikasi web sederhana untuk mengelola data karyawan menggunakan Go, MySQL, dan template HTML.

## Struktur Proyek

```
web_crud/
├── main.go                 # File utama untuk menjalankan aplikasi
├── controller/              # Folder untuk logika bisnis
│   ├── index_employee.go    # Controller untuk menampilkan daftar karyawan
│   ├── create_employee.go   # Controller untuk menambah karyawan baru
│   ├── update_employee.go   # Controller untuk mengupdate data karyawan
│   ├── delete_employee.go   # Controller untuk menghapus karyawan
│   └── hello_world.go      # Controller sederhana untuk testing
├── database/               # Folder untuk konfigurasi database
│   ├── database.go         # File koneksi database
│   └── schema.sql          # SQL script untuk membuat tabel
├── routes/                 # Folder untuk routing
│   └── router.go           # Konfigurasi URL endpoints
└── views/                  # Folder untuk template HTML
    ├── index.html          # Template untuk daftar karyawan
    ├── create.html         # Template untuk form tambah karyawan
    └── update.html         # Template untuk form update karyawan
```

## Cara Menjalankan

### 1. Persiapan Database

Pastikan MySQL sudah terinstall dan berjalan di localhost:3306 dengan user root (tanpa password).

Jalankan script SQL berikut untuk membuat database dan tabel:

```bash
mysql -u root < database/schema.sql
```

Atau jalankan perintah SQL secara manual dari file `database/schema.sql`.

### 2. Install Dependensi

```bash
go mod tidy
```

### 3. Jalankan Aplikasi

```bash
go run main.go
```

Aplikasi akan berjalan di http://localhost:8080

## Endpoint

| Method | URL                        | Deskripsi                         |
| ------ | -------------------------- | --------------------------------- |
| GET    | `/`                        | Menampilkan pesan "Hello World"   |
| GET    | `/employee`                | Menampilkan daftar semua karyawan |
| GET    | `/employee/create`         | Menampilkan form tambah karyawan  |
| POST   | `/employee/create`         | Menyimpan data karyawan baru      |
| GET    | `/employee/update?id={id}` | Menampilkan form update karyawan  |
| POST   | `/employee/update?id={id}` | Mengupdate data karyawan          |
| GET    | `/employee/delete?id={id}` | Menghapus data karyawan           |

## Fitur

1. **Read (Lihat Data)**

   - Menampilkan semua data karyawan dalam bentuk tabel
   - Setiap karyawan memiliki tombol Edit dan Hapus

2. **Create (Tambah Data)**

   - Form untuk menambah karyawan baru
   - Field: Nama, NPWP, Alamat
   - Validasi sederhana

3. **Update (Ubah Data)**

   - Form yang sudah terisi dengan data karyawan yang dipilih
   - Mengupdate data yang sudah ada

4. **Delete (Hapus Data)**
   - Menghapus karyawan berdasarkan ID
   - Konfirmasi sebelum menghapus

## Troubleshooting

### Data tidak muncul di /employee

1. Pastikan MySQL sudah berjalan
2. Verifikasi database dan tabel sudah dibuat menggunakan schema.sql
3. Periksa console output untuk pesan debug
4. Pastikan mengakses `/employee` (bukan `/employees`)

### Template parsing error

Jika ada error "ends in a non-text context":

- Pastikan textarea tidak memiliki newline sebelum template expression
- Gunakan `text/template` untuk update form jika perlu

### Error koneksi database

1. Pastikan MySQL berjalan di port 3306
2. Verifikasi user root ada dan tidak menggunakan password
3. Pastikan database `golang_crud` sudah dibuat

## Teknologi yang Digunakan

- **Go**: Bahasa pemrograman backend
- **MySQL**: Database untuk menyimpan data
- **HTML Template**: Rendering view di sisi server
- **Tailwind CSS**: Styling melalui CDN
- **HTTP ServeMux**: Routing sederhana di Go

## Kontribusi

1. Fork proyek ini
2. Buat branch fitur baru
3. Commit perubahan Anda
4. Push ke branch
5. Buat Pull Request

## Lisensi

MIT License
