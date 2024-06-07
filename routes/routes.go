package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/salwakhairu/Tugas-Besar-PBW-Sistem-Karyawan/controller"
	"golang.org/x/crypto/bcrypt"
)

// MapRoutes memetakan semua rute HTTP
func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/employee", controller.NewIndexEmployee(db))
	server.HandleFunc("/employee/create", controller.NewCreateEmployeeController(db))
	server.HandleFunc("/employee/update", controller.NewUpdateEmployeeController(db))
	server.HandleFunc("/employee/delete", controller.NewDeleteEmployeeController(db))
	server.HandleFunc("/login", LoginPage) // Gunakan fungsi LoginPage dari handlers
	server.HandleFunc("/login/submit", LoginHandler(db))
}

// LoginPage menampilkan halaman login
func LoginPage(w http.ResponseWriter, r *http.Request) {
	// Render halaman login di sini
	fp := filepath.Join("views", "login.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// LoginHandler memproses form login
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		var hashedPassword string
		err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hashedPassword)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Redirect ke halaman /employee setelah login berhasil
		http.Redirect(w, r, "/employee", http.StatusSeeOther)
	}
}
