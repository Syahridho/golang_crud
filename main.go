package main

import (
	"net/http"

	"github.com/syahridho/golang_crud/database"
	"github.com/syahridho/golang_crud/routes"
)

// Fungsi utama yang akan dijalankan pertama kali saat program start
func main() {
	// Inisialisasi koneksi ke database
	db := database.InitDatabase()
	
	// Membuat server HTTP multiplexer untuk menangani berbagai rute/endpoint
	server := http.NewServeMux()

	// Memetakan semua rute (URL) ke fungsi handler yang sesuai
	routes.MapRoutes(server,db)

	// Menjalankan web server di port 8080
	// Server akan berjalan terus-menerus menunggu request dari client
	http.ListenAndServe(":8080", server)
}