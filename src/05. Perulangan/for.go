package main

import "fmt"

func main() {

	fmt.Println("Perulangan dengan [Langsung mendeklarasikan kondisi awal, kondisi akhir, dan increment/decrement]")
	for  i := 0  ; i < 6 ; i++ {
		fmt.Printf("Perulangan ke-%d \n", i)
	}

	fmt.Println("Perulangan dengan [Deklarasi kondisi awal menggunakan variabel, kondisi akhir, dan increment/decrement di luar for]")
	var j = 1
	for j <= 5 {
		fmt.Println("Ini Angka", j)
		j++
	}

	fmt.Println("Perulangan dengan [Menggunakan for tanpa kondisi, menggunakn if (break) untuk menghentikan perulangan]")
	k := 0
	for {
		fmt.Println("Perulangan ke-", k)
		k++
		if k > 5 {
			break
		}
	}
}