package models

type Product struct {
	Id        string `json:"id"`
	Nama      string `json:"nama"`
	Harga     string `json:"harga"`
	Deskripsi string `json:"deskripsi"`
}
