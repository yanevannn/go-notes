package main

import (
	"fmt"
	"strings"
)

// Funsi yang menerima parameter berupa function
// Function ini akan menerima sebuah string dan sebuah function yang akan digunakan untuk memformat string tersebut
// Parameter kedua adalah function yang menerima string dan mengembalikan string
func formatHuruf (text string, formatText func(string) string){
	Hasil := formatText(text)
	fmt.Println(Hasil)
}

// Fungsi yang nantinya akan digunakan sebagai parameter pada function formatHuruf
func hurufBesar (text string) string {
	return strings.ToUpper(text)
}
// Fungsi yang nantinya akan digunakan sebagai parameter pada function formatHuruf
func hruufKecil (text string) string {
	return strings.ToLower(text)
}

// ===============================================
// Emnggunakan type untuk membuat function sebagai parameter yang bertujuan untuk memperjelas kode
// ===============================================
type FormatFunction func(string) string

func formatHuruf2 (text string, formatText FormatFunction){
	Hasil := formatText(text)
	fmt.Println(Hasil)
}

func main() {
	// Memanggil function formatHuruf dengan parameter string dan function hurufBesar
	formatHuruf("budi", hurufBesar)
	formatHuruf2("rudi", hurufBesar)
	// Memanggil function formatHuruf dengan parameter string dan function hruufKecil
	formatHuruf("BUDI", hruufKecil)
	formatHuruf2("RUDI", hruufKecil)
}