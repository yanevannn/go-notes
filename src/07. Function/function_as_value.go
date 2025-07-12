package main

import "fmt"

func pertambahan(a, b int) int {
	return a + b 
}

func main () {
	// Function sebagai value
	// Kita bisa menyimpan function ke dalam variabel
	// Sehingga kita bisa memanggil function tersebut melalui variabel
	hasilPertambahan := pertambahan

	fmt.Println("Hasil pertambahan 10 + 20 =", hasilPertambahan(10, 20)) // Memanggil function melalui variabel
}