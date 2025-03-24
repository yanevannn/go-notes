package main

import "fmt"

func main() {
    // Deklarasi konstanta
    const pi = 3.14
    const nama = "Evan"
    const umur = 20

    // Menggunakan konstanta dalam program
    fmt.Println("Nilai Pi:", pi)
    fmt.Println("Nama Saya:", nama)
    fmt.Println("Umur Saya:", umur)

    // Contoh dengan operasi matematika
    const jariJari = 5
    var luasLingkaran = pi * jariJari * jariJari

    fmt.Println("Luas Lingkaran apabila jari-jari:",jariJari,"adalah", luasLingkaran)
}
