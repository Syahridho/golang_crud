package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// Employee struct untuk menyimpan data employee dari database
type Employee struct {
	Id string      // ID employee (primary key)
	Name string    // Nama employee
	NPWP string   // Nomor NPWP employee
	Address string // Alamat employee
}

// NewIndexEmployee fungsi untuk membuat handler yang menampilkan daftar semua employee
// Parameter: db - koneksi database untuk mengambil data employee
// Mengembalikan fungsi handler untuk HTTP request
func NewIndexEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request)  {
	return func(w http.ResponseWriter, r *http.Request) {
		// Debug log untuk memeriksa apakah handler dipanggil
		fmt.Println("DEBUG: IndexEmployee handler called")
		
		// Query untuk mengambil semua data employee dari tabel
		rows, err := db.Query("SELECT id, name, npwp, address FROM employee")

		if err != nil {
			fmt.Printf("DEBUG: Database query error: %v\n", err)
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer rows.Close() // Pastikan rows ditutup setelah selesai digunakan

		var employees []Employee // Slice untuk menyimpan semua data employee

		// Loop melalui setiap baris hasil query
		for rows.Next() {
			var employee Employee // Variabel untuk menyimpan satu data employee

			// Scan data dari baris ke struct employee
			err = rows.Scan(
				&employee.Id,
				&employee.Name,
				&employee.NPWP,
				&employee.Address,
			)

			if err != nil {
				fmt.Printf("DEBUG: Row scan error: %v\n", err)
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Tambahkan employee ke slice
			employees = append(employees, employee)
		}
		
		fmt.Printf("DEBUG: Found %d employees\n", len(employees))

		// Path ke file template HTML
		fp := filepath.Join("views", "index.html")
		tmpl, err := template.ParseFiles(fp)

		if err != nil {
			fmt.Printf("DEBUG: Template parse error: %v\n", err)
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Buat map untuk data yang akan dikirim ke template
		data := make(map[string]any)
		data["employees"] = employees  // Key "employees" digunakan di template

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