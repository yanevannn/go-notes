package main

import "fmt"

func perhitunganRekursif(angka int) int {
	if angka == 3 {
		return angka
	}
	fmt.Println("Angka sekarang:", angka)
	return perhitunganRekursif(angka + 1)
}

//Contohnya perhitungan factorial dengan rekursif function
func factorial(n int) int{
	if n == 0{
		return 1
	}
	return n * factorial(n-1)
}


func main() {
	//constoh 1
	perhitunganRekursif(-2)
	
	//contoh 2 factorial
	nilai := 4
	factorialResult := factorial(nilai)
	fmt.Printf("Hasil factorial dari %d adalah %d\n", nilai, factorialResult)
}
