package main

import "fmt"

func main() {
	// Format Array Basic => var namaArray [jumlahElemen]tipeData
	var status [3]string // Deklarasi array dengan 3 elemen bertipe string
	status[0] = "Aktif"
	status[1] = "Non-Aktif"
	status[2] = "Pending"

	fmt.Println("Jumlah Data dalam array status:", len(status)) //Fungsi len() untuk mendapatkan jumlah elemen/data dalam array
	fmt.Println("Data Array status:", status)
	// Mengakses elemen array => namaArray[index]
	for i :=0; i < len(status); i++{
		fmt.Printf("Data array status ke- %d: %s \n",i, status[i])
	}

	// Format Array tanpa deklarasi jumlah elemen
	// var namaArray = [...]tipeData{data1, data2, data3, ...}
	var angka = [...]int{1,2,3,4,5}
	fmt.Println("Jumlah Data dalam array angka:", len(angka))
	fmt.Println("Data array angka:", angka)
}