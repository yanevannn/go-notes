package main

import "fmt"

func main() {
	// Deklarasi Variabel
	var namaDepan string = "Yan"
	var namaBelakang string
	namaBelakang = "Evan"

	fmt.Println("Nama Lengkap:", namaDepan, namaBelakang)

	// Deklarasi Variabel Tanpa Tipe Data
	var alamatRumah = "Denpasar"
	alamarKantor := "Renon"
	fmt.Println("Alamat Rumah: " + alamatRumah)
	fmt.Println("Alamat Kantor: " + alamarKantor)

	// Deklarasi Variabel Underscore 
	// Underscore digunakan untuk variabel yang tidak digunakan & Biasanya digunakan untuk menghindari error jika ada variabel yang tidak digunakan
	var email, _ = "evan@mail.com", "Ini variabel yang tidak digunakan"
	fmt.Println("alamat email:", email)
}