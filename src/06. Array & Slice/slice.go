package main

import "fmt"

func main() {
	//Slice adalah tipe data yang digunakan untuk menyimpan kumpulan data yang dinamis
	//Slice adalah bagian dari array yang dapat berubah ukurannya
	//Slice tidak memiliki ukuran tetap berbeda dengan array yang memiliki ukuran tetap

	var buah []string
	fmt.Println("[Sebelum ditambahkan data] Jumlah Data Slice buah:", len(buah))

	// Menambahkan data ke dalam slice menggunakan fungsi append()
	// Fungsi append() digunakan untuk menambahkan data baru ke dalam slice dimana ukuran slice akan bertambah secara otomatis
	// Append dapat menambahkan lebih dari satu data sekaligus
	buah = append(buah, "Jeruk", "Apel", "Mangga", "Pisang", "Semangka")
	fmt.Println("[Sesudah ditambahkan data] Jumlah Data Slice buah:", len(buah))
	fmt.Println("Data Slice Buah:", buah)
}