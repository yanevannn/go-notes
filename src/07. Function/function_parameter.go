package main

import "fmt"

// Function dengan parameter
func perkenalanDiri (nama string, umur int){
	fmt.Println("Halo perkenalkan nama saya", nama)
	fmt.Printf("Umur saya %d tahun \n", umur)
}

func main() {
	perkenalanDiri("Evan", 20)
	perkenalanDiri("Gede", 17)
}