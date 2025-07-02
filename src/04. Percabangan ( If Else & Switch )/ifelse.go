package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukkan nilai (0-100): ")
	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100{
		fmt.Println("Nilai A")
	} else if nilai >= 80 && nilai < 90 {
		fmt.Println("Nilai B")
	} else if nilai >= 70 && nilai < 80 {
		fmt.Println("Nilai C")
	} else if nilai >= 60 && nilai < 70 {
		fmt.Println("Nilai D")
	} else if nilai >= 0 && nilai < 60 {
		fmt.Println("Nilai E")
	} else {
		fmt.Println("Nilai tidak valid, masukkan antara 0-100")
	}
}