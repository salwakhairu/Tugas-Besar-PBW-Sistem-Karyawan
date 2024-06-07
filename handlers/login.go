package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

// LoginPageHandler menampilkan halaman login
func LoginPage(w http.ResponseWriter, r *http.Request) {
	fp := filepath.Join("views", "login.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
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
			fmt.Println("Error retrieving hashed password from database:", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			fmt.Println("Incorrect password:", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Login berhasil, redirect ke halaman lain atau set session
		http.Redirect(w, r, "/employee", http.StatusSeeOther)
	}
}
