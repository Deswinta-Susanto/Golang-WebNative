package config

import (
	"database/sql"
    "log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// ConnectDB menghubungkan aplikasi dengan database MySQL
func ConnectDB() *sql.DB {
	// Membuka koneksi ke database
	db, err := sql.Open("mysql", "root:new_password@/go_db?parseTime=true")
	if err != nil {
		panic(err) // Jika terjadi error, program akan berhenti
	}

	// Pastikan koneksi berhasil
	err = db.Ping()
	if err != nil {
		panic(err) // Jika gagal ping database, program akan berhenti
	}
    log.Println("Database connected")

	// Menyimpan koneksi ke variabel global DB
	DB = db

	// Mengembalikan koneksi DB untuk digunakan
	return DB
}
