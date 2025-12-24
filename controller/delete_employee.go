package controller

import (
	"database/sql"
	"fmt"
	"net/http"
)

// NewDeleteEmployeeController fungsi untuk membuat handler yang menangani penghapusan employee
// Parameter: db - koneksi database untuk menghapus data employee
// Mengembalikan fungsi handler untuk HTTP request
func NewDeleteEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request)  {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ambil ID employee dari URL query parameter
		id := r.URL.Query().Get("id")
		fmt.Printf("DEBUG: Loading employee with ID: %s\n", id)

		// Hapus data employee dari database menggunakan parameterized query
		_, err := db.Exec("DELETE FROM employee WHERE id = ?", id)

		if err != nil {
			fmt.Printf("DEBUG: Query error: %v\n", err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		// Redirect ke halaman daftar employee setelah berhasil menghapus
		http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
	}
		
}