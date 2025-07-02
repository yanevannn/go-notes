package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan angka (1-10): ")
	fmt.Scan(&angka)
	switch angka {
	case 1, 3, 5, 7, 9:
		fmt.Println("Angka", angka, "adalah bilangan ganjil")
	case 2, 4, 6, 8, 10:
		fmt.Println("Angka", angka, "adalah bilangan genap")
	default:
		fmt.Println("Angka tidak valid, masukkan antara 1-10")
	}
}
