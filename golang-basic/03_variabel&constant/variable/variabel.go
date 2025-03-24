package main

import "fmt"

func main() {
	// 1. Deklarasi variabel tanpa inisialisasi (nilai default akan diberikan)
    var angka1 int
    var kata string
    var status bool

    fmt.Println("angka1:", angka1)   // Output: 0 (default int)
    fmt.Println("kata:", kata)       // Output: "" (string kosong)
    fmt.Println("status:", status)   // Output: false (default bool)

    // 2. Deklarasi variabel dengan inisialisasi langsung
    var angka2 int = 10
    var nama string = "Golang"
    var isActive bool = true

    fmt.Println("angka2:", angka2)
    fmt.Println("nama:", nama)
    fmt.Println("isActive:", isActive)

    // 3. Deklarasi dengan tipe data otomatis (tanpa menulis tipe data)
    var angka3 = 20.5
    var pesan = "Hello, World!"

    fmt.Println("angka3:", angka3) // Tipe data float64
    fmt.Println("pesan:", pesan)   // Tipe data string

    // 4. Deklarasi multiple variabel dalam satu baris
    var x, y, z int = 1, 2, 3
    fmt.Println("x:", x, "y:", y, "z:", z)

    var a, b, c = 4.5, "Variabel", true
    fmt.Println("a:", a, "b:", b, "c:", c)

    // 5. Variabel dalam blok deklarasi
    var (
        umur  int    = 25
        kota  string = "Jakarta"
        aktif bool   = true
    )
    fmt.Println("Umur:", umur, "Kota:", kota, "Aktif:", aktif)
}