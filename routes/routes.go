package handlers

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// LoginPage menampilkan halaman login
func LoginPage(w http.ResponseWriter, r *http.Request) {
	// Logika untuk menampilkan halaman login
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
