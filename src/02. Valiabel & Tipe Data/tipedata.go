package main

import "fmt"

func main() {
	// Tipe data Integer (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64)
	var umur int = 20
	var gaji int64 = 1000000
	fmt.Print("Umur: ", umur, "\n")
	fmt.Print("Gaji: ", gaji, "\n")

	// Tipe data Float (float32, float64)
	var tinggi float32 = 1.75
	var berat float64 = 70.5
	fmt.Printf("Tinggi: %.2f m\n", tinggi)
	fmt.Printf("Berat: %.2f kg\n", berat)

	// Tipe data String (Untuk teks)
	var nama string = "Yan Evan"
	fmt.Println("Nama:", nama)

	// Tipe data Boolean (true/false - untuk kondisi logika, ditulis dengan huruf kecil)
	var isMarried bool = true
	fmt.Printf("Apakah sudah menikah? %t\n", isMarried) 
	fmt.Println("Apakah sudah menikah?", isMarried)

}