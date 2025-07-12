package main

import "fmt"

// Function variadic adalah function yang dapat menerima jumlah parameter yang tidak terbatas
// Parameter variadic ditandai dengan ... sebelum tipe data parameter
func penjumlahan(angka ...int) int {
	total := 0
	for _, a := range angka{
		total += a
	}
	return total
}

func main() {
	// Contoh penggunaan function variadic
	fmt.Println("Hasil penjumlahan 1 + 2 + 3 =", penjumlahan(1, 2, 3))
	
	// Kita juga bisa mengirimkan slice ke function variadic
	nilai := []int{10, 20, 30}
	fmt.Println("Hasil penjumlahan dari slice =", penjumlahan(nilai...)) // Perhatikan penggunaan ... untuk mengirimkan slice sebagai variadic argument
}