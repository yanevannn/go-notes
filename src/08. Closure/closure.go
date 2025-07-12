package main

import "fmt"

func main() {
	// Variabel dalam scope fungsi main
	nilai_a := 0
	nilai_b := 10

	fmt.Println("Nilai awal 'nilai_a' di scope main:", nilai_a)
	fmt.Println("Nilai awal 'nilai_b' di scope main:", nilai_b)

	// Closure yang mengakses dan memodifikasi variabel dari scope luar
	rubahAngka := func() {
		// Mengakses dan mengubah variabel 'nilai_a' dari scope fungsi main
		// Karena tidak dibuat ulang di dalam closure, perubahan ini akan memengaruhi nilai asli di scope luar
		nilai_a = 20
		fmt.Println("Nilai 'nilai_a' di scope main setelah diubah dalam closure:", nilai_a)

		// Membuat variabel baru dengan nama yang sama 'nilai_b' di dalam closure (shadowing)
		// Variabel ini hanya hidup di dalam closure, sehingga tidak memengaruhi 'nilai_b' yang ada di scope fungsi main
		nilai_b := 30
		fmt.Println("Variabel 'nilai_b' baru dalam scope closure:", nilai_b)
	}

	// Menjalankan closure
	rubahAngka()

	// Mengecek kembali nilai setelah closure dijalankan
	fmt.Println("Nilai akhir 'nilai_a' di scope main:", nilai_a)
	fmt.Println("Nilai akhir 'nilai_b' di scope main (tidak berubah):", nilai_b)
}
