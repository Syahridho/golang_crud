package routes

import (
	"database/sql"
	"net/http"

	"github.com/syahridho/golang_crud/controller"
)

// MapRoutes fungsi untuk memetakan URL endpoint ke fungsi controller yang sesuai
// Parameter:
// - server: HTTP ServeMux untuk menangani routing
// - db: Koneksi database yang akan digunakan oleh controller
func MapRoutes(server *http.ServeMux, db *sql.DB) {
	// Route untuk halaman utama, menampilkan "Hello World"
	server.HandleFunc("/", controller.NewHelloWorldController())
	
	// Route untuk menampilkan daftar semua employee
	server.HandleFunc("/employee", controller.NewIndexEmployee(db))
	
	// Route untuk membuat employee baru (GET untuk form, POST untuk submit)
	server.HandleFunc("/employee/create", controller.NewCreateEmployeeController(db))
	
	// Route untuk mengupdate employee yang sudah ada (GET untuk form, POST untuk submit)
	server.HandleFunc("/employee/update", controller.NewUpdateEmployeeController(db))
	
	// Route untuk menghapus employee (GET/POST untuk konfirmasi dan eksekusi)
	server.HandleFunc("/employee/delete", controller.NewDeleteEmployeeController(db))
}