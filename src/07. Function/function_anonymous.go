package main

import "fmt"

func main() {
	// Anonymous function dibuat didalam variabel [Sama seperti arrow function pada javascript]
	getHello := func() {
		fmt.Println("Hello Evan")
	}
	// Memanggil anonymous function
	getHello()


	// Anonymous function dapat disii parameter seperti function biasa
	getHelloWithName := func(name string) {
		fmt.Println("Hello", name)
	}
	getHelloWithName("Made")

	// Anonymous function dapat juga diisi parameter dan return value
	perhitanganTambah := func(a int, b int) int{
		return a + b
	}
	fmt.Println("Hasil pertambhan 10 + 20 =", perhitanganTambah(10, 20))
}
