package models

type Product struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Nama  string `json:"nama"`
	Stok  int    `json:"stok"`
	Harga int    `json:"harga"`
}
