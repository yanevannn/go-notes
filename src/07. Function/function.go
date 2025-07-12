package main

import "fmt"

// Apabila penamaan function menggunakan huruf kecil, maka function tersebut hanya bisa diakses di dalam package yang sama.
// Apabila penamaan function menggunakan huruf besar, maka function tersebut bisa diakses dari package lain.

func greatingMessage (){
	fmt.Println("Hello Semua, Selamat Datang di Go Lang")
}

func main() {
	greatingMessage()
	greatingMessage()
	greatingMessage()
}