package main

import (
	"net/http"

	"github.com/salwakhairu/Tugas-Besar-PBW-Sistem-Karyawan/database"
	"github.com/salwakhairu/Tugas-Besar-PBW-Sistem-Karyawan/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}
