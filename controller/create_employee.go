package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

// NewCreateEmployeeController fungsi untuk membuat handler yang menangani pembuatan employee baru
// Parameter: db - koneksi database untuk menyimpan data employee
// Mengembalikan fungsi handler untuk HTTP request
func NewCreateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request)  {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Parse form data dari request
			r.ParseForm()

			// Ambil data dari form
	 		name := r.Form["name"][0]      // Nama employee
			npwp := r.Form["npwp"][0]      // NPWP employee
			address := r.Form["address"][0]  // Alamat employee

			// Insert data employee ke database menggunakan parameterized query untuk mencegah SQL injection
			_, err := db.Exec("INSERT INTO employee (name, npwp, address) VALUES (?, ?, ?)", name, npwp, address)

			if err != nil {
				w.Write([]byte(err.Error()))
 				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			
			// Redirect ke halaman daftar employee setelah berhasil menyimpan
			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			// Untuk request GET, tampilkan form pembuatan employee
			
			// Path ke file template HTML
			fp := filepath.Join("views", "create.html")
			tmpl, err := template.ParseFiles(fp)
	
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
	
			// Eksekusi template tanpa data (form kosong)
			err = tmpl.Execute(w, nil)
	
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
		
}