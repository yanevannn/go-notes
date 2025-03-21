# Pengenalan Golang

## 🚀 Apa itu Golang?
Go (Golang) adalah bahasa pemrograman yang dibuat oleh Google. Bahasa ini mirip dengan C tetapi lebih sederhana dan modern.  

Contoh program Go:
```go
package main

func main() {
  // Kode program ditulis di sini
}
```

## 🛠 Persiapan  

### 📥 Instalasi Go  
1. **Download** Go dari [situs resminya](https://golang.org/doc/install).  
2. **Install** sesuai sistem operasi:  
   - **Windows**: Jalankan file installer, lalu periksa dengan:  
     ```sh
     go version
     ```
   - **Mac (Menggunakan Homebrew)**:  
     ```sh
     brew install go
     ```
     Setelah instalasi, periksa dengan:  
     ```sh
     go version
     ``` 
   - **Mac/Linux (Manual)**: Ekstrak file dan tambahkan Go ke `PATH`.  

### ▶️ Menjalankan Program  
Program Go dapat dijalankan dengan 2 cara yaitu :
1. **Kompilasi & Jalankan**:  
   ```sh
   go build main.go  
   ./main  
   ```
2. **Jalankan langsung tanpa kompilasi**:  
   ```sh
   go run main.go  
   ```

## 📂 Struktur Program Dasar  
Setiap program Go memiliki bagian utama seperti:  
```go
package main  // Wajib ada di setiap program utama

import "fmt"  // Menggunakan package untuk output teks

func main() {  
  fmt.Println("Hello, World!")  // Cetak teks ke layar
}
```

### 🔹 Package  
Semua kode di Go harus berada dalam **package**. Program utama selalu menggunakan:  
```go
package main
```

### 🔹 Import  
Untuk menggunakan fungsi dari package lain, kita mengimpornya:  
```go
import "fmt"
```
Kita juga bisa mengimpor beberapa package sekaligus:  
```go
import (
  "fmt"
  "os"
)
```

### 🔹 Fungsi `main()`  
Fungsi ini wajib ada karena akan dieksekusi pertama kali:  
```go
func main() {
  fmt.Println("Halo, dunia!")
}
```

## ✍️ Menulis Program Pertama  
Buat file baru bernama `hello.go`, lalu ketik:  
```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```
Jalankan dengan:  
```sh
go run hello.go
```

✅ **Selamat!** Kamu sudah membuat program pertamamu di Go. 🎉

