package main

import "fmt"

func main() {
	// Konstanta (Nialai yang tidak bisa diubah)
	const pi float64 = 3.14
	fmt.Println("Nilai Pi:", pi)

	// Operator Aritmatika (tambah [+], kurang [-], kali [*], bagi [/], modulus [%])
	var a int = 10
	var b int = 3
	fmt.Printf("Penjumlahan: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Pengurangan: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Perkalian: %d x %d = %d\n", a, b, a*b)
	fmt.Printf("Pembagian: %d / %d = %.2f\n", a, b, float64(a)/float64(b))
	fmt.Printf("Modulus: Hasil sisa bagi dari %d / %d = %d\n", a, b, a%b)

	// Operator Perbandingan // (Operator ini menghasilkan nilai boolean true/false)
	// (sama dengan [==]
	fmt.Println("Apakah a sama dengan b?", a == b) // false
	// tidak sama dengan [!=]
	fmt.Println("Apakah a tidak sama dengan b?", a != b) // true
	// lebih besar dari [>]
	fmt.Println("Apakah a lebih besar dari b?", a > b) // true
	// lebih kecil dari [<]
	fmt.Println("Apakah a lebih kecil dari b?", a < b) // false
	// lebih besar atau sama dengan [>=]
	fmt.Println("Apakah a lebih besar atau sama dengan b?", a >= b) // true
	// lebih kecil atau sama dengan [<=])
	fmt.Println("Apakah a lebih kecil atau sama dengan b?", a <= b) // false


	// Operator Logika (dan [&&], atau [||], tidak [!])
	fmt.Println("Apakah a lebih besar dari 5 dan b lebih kecil dari 5?", a > 5 && b < 5) // true
	fmt.Println("Apakah a lebih besar dari 5 atau b lebih kecil dari 5?", a > 5 || b < 5) // true
	fmt.Println("Apakah bukan a lebih besar dari 5?", !(a > 5)) // false
}