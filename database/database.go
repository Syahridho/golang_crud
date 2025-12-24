package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// InitDatabase fungsi untuk menginisialisasi koneksi ke database MySQL
// Mengembalikan objek database yang siap digunakan
func InitDatabase() *sql.DB {
	fmt.Println("DEBUG: Initializing database connection...")

	// DSN (Data Source Name) untuk koneksi ke MySQL
	// Format: user@tcp(host:port)/database_name
	dsn := "root@tcp(localhost:3306)/golang_crud"
	fmt.Printf("DEBUG: Connecting with DSN: %s\n", dsn)

	// Membuka koneksi ke database menggunakan driver MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("DEBUG: Error opening database: %v\n", err)
		panic(err) // Hentikan program jika koneksi gagal
	}

	// Test koneksi dengan ping ke database
	err = db.Ping()
	if err != nil {
		fmt.Printf("DEBUG: Error pinging database: %v\n", err)
		panic(err) // Hentikan program jika ping gagal
	}

	fmt.Println("DEBUG: Database connection successful")
	
	// Periksa apakah tabel employee ada dan memiliki data
	checkTable(db)
	
	return db // Kembalikan objek database yang sudah terkoneksi
}

// checkTable fungsi untuk memeriksa keberadaan tabel employee dan jumlah datanya
func checkTable(db *sql.DB) {
	fmt.Println("DEBUG: Checking if employee table exists...")
	
	// Query untuk memeriksa apakah tabel employee ada
	rows, err := db.Query("SHOW TABLES LIKE 'employee'")
	if err != nil {
		fmt.Printf("DEBUG: Error checking table: %v\n", err)
		return
	}
	defer rows.Close() // Pastikan rows ditutup setelah selesai
	
	if !rows.Next() {
		fmt.Println("DEBUG: Employee table does not exist!")
		return
	}
	
	fmt.Println("DEBUG: Employee table exists")
	
	// Hitung jumlah data dalam tabel employee
	countRows, err := db.Query("SELECT COUNT(*) FROM employee")
	if err != nil {
		fmt.Printf("DEBUG: Error counting employees: %v\n", err)
		return
	}
	defer countRows.Close() // Pastikan countRows ditutup setelah selesai
	
	var count int
	if countRows.Next() {
		countRows.Scan(&count) // Ambil hasil query ke variabel count
		fmt.Printf("DEBUG: Found %d employees in the table\n", count)
	}
}