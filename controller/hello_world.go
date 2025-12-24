package controller

import "net/http"

// NewHelloWorldController fungsi untuk membuat handler yang menampilkan "Hello World"
// Tidak memerlukan parameter database karena hanya menampilkan teks statis
// Mengembalikan fungsi handler untuk HTTP request
func NewHelloWorldController() func(w http.ResponseWriter, r *http.Request)  {
	return func(w http.ResponseWriter, r *http.Request) {
		// Kirim response "Hello World" ke client
		w.Write([]byte("Hello World"))
	}
}