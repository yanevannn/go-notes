package main

import "fmt"

// Function dengan return value
func perhituganKuadrat(angka int) int {
	return angka * angka
}

// Function dengan named return value [Kita bisamemberikan nama pada return value]
func perhitunganKubik (angka int) (hasil int){
	hasil = angka * angka * angka
	return // Kita bisa langsung return tanpa menyebutkan nama variabelnya karena sudah di definikan sebelumnya
}

// Function dengan Multiple return value
func perhitunganKuadratKubik (angka int) (kuadrat int, kubik int){
	kuadrat = angka * angka
	kubik = angka * angka * angka
	return // Kita bisa langsung return tanpa menyebutkan nama variabelnya karena sudah di definikan sebelumnya
}

func main() {
	var nilai int = 10
	fmt.Printf("Hasil %d^2 adalah %d\n", nilai, perhituganKuadrat(nilai) )

	// Menyimpan nilai return function ke dalam variabel
	hasilKubik := perhitunganKubik(nilai)
	fmt.Printf("Hasil %d^3 adalah %d\n", nilai, hasilKubik )

	// Menggunakan Multiple return value
	nilaiKuadrat, nilaiKubik := perhitunganKuadratKubik(nilai)
	fmt.Printf("%d^2 = %d dan %d^3 = %d \n", nilai, nilaiKuadrat, nilai, nilaiKubik)

	// Ignore return value menggunakan _ (underscroe) -> hal ini berguna jika kita hanya ingin mengambil salah satu nilai return value
	contoh := 5
	_, nilaiKubikContoh := perhitunganKuadratKubik(contoh) // Contoh Kita hanya mengambil nilai kubik saja
	fmt.Printf("Hasil %d^3 adalah %d\n", contoh, nilaiKubikContoh)
}