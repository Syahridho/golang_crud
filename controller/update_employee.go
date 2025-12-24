package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"
	texttmpl "text/template" // Menggunakan text/template untuk menghindari auto-escaping HTML
)

// NewUpdateEmployeeController fungsi untuk membuat handler yang menangani update data employee
// Parameter: db - koneksi database untuk mengupdate data employee
// Mengembalikan fungsi handler untuk HTTP request
func NewUpdateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request)  {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Parse form data dari request
			r.ParseForm()
			
			// Ambil ID employee dari URL query parameter
			id := r.URL.Query().Get("id")
			// Ambil data dari form
	 		name := r.Form["name"][0]      // Nama employee yang diupdate
			npwp := r.Form["npwp"][0]      // NPWP employee yang diupdate
			address := r.Form["address"][0]  // Alamat employee yang diupdate

			fmt.Printf("DEBUG: Updating employee ID %s with name: %s\n", id, name)
			
			// Update data employee di database menggunakan parameterized query
			_, err := db.Exec("UPDATE employee SET name = ?, npwp = ?, address = ? WHERE id = ?", name, npwp, address, id)

			if err != nil {
				fmt.Printf("DEBUG: Update error: %v\n", err)
				w.Write([]byte(err.Error()))
 				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			
			fmt.Println("DEBUG: Update successful, redirecting to /employee")
			// Redirect ke halaman daftar employee setelah berhasil update
			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			// Untuk request GET, tampilkan form update dengan data employee yang ada
			
			// Ambil ID employee dari URL query parameter
			id := r.URL.Query().Get("id")
			fmt.Printf("DEBUG: Loading employee with ID: %s\n", id)

			// Query untuk mengambil data employee berdasarkan ID
			row := db.QueryRow("SELECT name, npwp, address FROM employee WHERE id = ?", id)

			if row.Err() != nil {
				fmt.Printf("DEBUG: Query error: %v\n", row.Err())
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var employee Employee // Variabel untuk menyimpan data employee

			// Scan data dari query ke struct employee
			err := row.Scan(
				&employee.Name,
				&employee.NPWP,
				&employee.Address,
			)
			employee.Id = id // Set ID employee

			if err != nil {
				fmt.Printf("DEBUG: Scan error: %v\n", err)
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			
			// Path ke file template HTML
			fp := filepath.Join("views", "update.html")
			// Parse template menggunakan text/template (bukan html/template)
			tmpl, err := texttmpl.ParseFiles(fp)
	
			if err != nil {
				fmt.Printf("DEBUG: Template error: %v\n", err)
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
	
			// Buat map untuk data yang akan dikirim ke template
			data := make(map[string]any)
			data["employee"] = employee // Key "employee" digunakan di template

			// Eksekusi template dengan data employee
			err = tmpl.Execute(w, data)
			if err != nil {
				fmt.Printf("DEBUG: Template execution error: %v\n", err)
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
		
}