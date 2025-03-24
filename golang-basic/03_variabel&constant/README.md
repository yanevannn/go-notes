# Variabel dan Konstanta dalam Golang

## Variabel
Variabel adalah tempat untuk menyimpan data yang bisa diubah selama program berjalan. Setiap variabel memiliki tipe data tertentu yang menentukan cara penyimpanannya dan jenis operasi yang bisa dilakukan.

### **Aturan Penamaan Variabel**
- Bisa terdiri dari huruf, angka, dan garis bawah (`_`).
- Harus dimulai dengan huruf atau garis bawah.
- Go peka terhadap huruf besar-kecil (misalnya, `nama` dan `Nama` dianggap berbeda).

### **Cara Mendeklarasikan Variabel**
Format dasar untuk mendeklarasikan variabel:
```go
var nama_variabel tipe_data;
```
Contoh:
```go
var angka1, angka2, angka3 int;
var huruf1, huruf2 byte;
var tinggi, berat float32;
```

### **Memberikan Nilai ke Variabel**
Kita bisa langsung memberikan nilai saat mendeklarasikan variabel:
```go
nama_variabel = nilai;
```
Contoh:
```go
var angka int = 32;

//atau secara otomatis

var x = 3;
```

### **Variabel dengan Tipe Data Tetap**
Variabel bisa dideklarasikan dengan tipe data tetap, sehingga hanya bisa diisi dengan nilai yang sesuai (semisal sebuah variabel sudah dideklarasikan dengan float namun diisi dengan string maka akan menghasilkan eror):
```go
package main

import "fmt"

func main() {
  var angka float64
  angka = 42.5  
  fmt.Println(angka)
  fmt.Printf("Tipe variabel angka adalah %T\n", angka)
}
```

### **Variabel dengan Tipe Data Otomatis**
Go bisa langsung menentukan tipe data berdasarkan nilai yang diberikan:
```go
package main

import "fmt"

func main(){
  var angka_saya float64 = 29.9
  angka_lain := 32  // tanda := digunakan untuk deklarasi otomatis
  
  fmt.Println(angka_saya)
  fmt.Println(angka_lain)

  fmt.Printf("Tipe data angka_saya adalah %T\n", angka_saya)
  fmt.Printf("Tipe data angka_lain adalah %T\n", angka_lain)
}
```

### **Deklarasi Variabel Beragam Tipe Sekaligus**
Kita bisa mendeklarasikan beberapa variabel dengan tipe berbeda dalam satu baris:
```go
package main

import "fmt"

func main() {
   var a, b, c = 3, 4, "bellshade"  
   
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   fmt.Printf("Tipe a: %T\n", a)
   fmt.Printf("Tipe b: %T\n", b)
   fmt.Printf("Tipe c: %T\n", c)
}
```

---

## **Konstanta**
Konstanta adalah nilai tetap yang tidak bisa diubah setelah dideklarasikan. Contoh konstanta adalah angka `3.14` untuk nilai `pi`.

### **Cara Menuliskan Konstanta**
Konstanta bisa berupa bilangan bulat, desimal, atau heksadesimal:
```go
const pi = 3.14
const angka_hex = 0x4b
const angka_octal = 0123
```

### **Karakter Escape**

| Escape String | Arti | Contoh | Hasil |
|--------------|------|--------|-------|
| `\\` | Karakter `\` | `fmt.Println("\\")` | `\` |
| `\'` | Karakter `'` | `fmt.Println("\'")` | `'` |
| `\"` | Karakter `"` | `fmt.Println("\"")` | `"` |
| `\A` | Alert/Peringatan | - | (tidak umum digunakan) |
| `\B` | Menghapus | - | (tidak umum digunakan) |
| `\F` | Form Feed | - | (jarang digunakan) |
| `\n` | Baris baru | `fmt.Println("Hello\nWorld")` | `Hello`<br>`World` |
| `\R` | Carriage Return | - | (jarang digunakan) |
| `\t` | Indentasi horizontal (Tab) | `fmt.Println("Hello\tWorld")` | `Hello    World` |
| `\v` | Indentasi vertikal | - | (jarang digunakan) |
| `\ooo` | Bilangan oktal (1-3 digit) | - | (tergantung nilai oktal) |
| `\xhh...` | Bilangan heksadesimal | - | (tergantung nilai hex) |

Contoh penggunaan escape string `\t`:
```go
package main

import "fmt"

func main() {
    fmt.Println("bell\thade")
}
```
Hasilnya:
```
bell    hade
```
---

## **Mendeklarasikan Konstanta dengan `const`**
Kata kunci `const` digunakan untuk mendeklarasikan konstanta:
```go
const nama_variabel tipe_data = nilai;
```
Contoh penggunaan:
```go
package main

import "fmt"

func main(){
    const PANJANG = 12
    const LEBAR = 2

    var hasil int
    hasil = PANJANG * LEBAR

    fmt.Printf("Hasil: %d", hasil)
}
```

---

## **Perbedaan Variabel dan Konstanta**
| Perbedaan | Variabel (`var`) | Konstanta (`const`) |
|-----------|-----------------|---------------------|
| **Nilai** | Bisa diubah | Tidak bisa diubah |
| **Harus diinisialisasi?** | Tidak wajib langsung diisi | Harus langsung diisi |
| **Bisa pakai `:=`?** | Ya | Tidak |
| **Digunakan untuk** | Data yang bisa berubah | Data yang tetap sepanjang program |

Gunakan `const` jika nilainya tidak akan berubah, dan `var` jika nilainya bisa berubah selama program berjalan.

