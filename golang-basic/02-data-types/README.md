

# Tipe Data dalam Golang  

Golang memiliki beberapa tipe data utama, yaitu **String**, **Number**, dan **Boolean**. Setiap tipe data memiliki karakteristik dan kegunaannya masing-masing.  

## Tipe Data String  
Tipe data **String** adalah kumpulan karakter yang jumlahnya bisa **nol** hingga **tak hingga**. String direpresentasikan dengan **kata kunci `string`**.  
Nilai string harus diawali dan diakhiri dengan **tanda petik dua (`""`)**.  

**Contoh:**  
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Golang!") // Output: Hello, Golang!
}
```

---

## Tipe Data Number  
Golang memiliki dua jenis tipe data numerik, yaitu **Integer** dan **Floating Point**.  

### 🔹 Integer  
Tipe **integer** digunakan untuk menyimpan angka **bilangan bulat**.  

| Jenis  | Nilai Minimum | Nilai Maksimum |
|--------|--------------|--------------|
| `int8`  | -128 | 127 |
| `int16` | -32,768 | 32,767 |
| `int32` | -2,147,483,648 | 2,147,483,647 |
| `int64` | -9,223,372,036,854,775,808 | 9,223,372,036,854,775,807 |

Selain `int`, terdapat juga **unsigned integer (`uint`)**, yang hanya menyimpan angka **positif**.  

| Jenis  | Nilai Minimum | Nilai Maksimum |
|--------|--------------|--------------|
| `uint8`  | 0 | 255 |
| `uint16` | 0 | 65,535 |
| `uint32` | 0 | 4,294,967,295 |
| `uint64` | 0 | 18,446,744,073,709,551,615 |

**Contoh:**  
```go
package main

import "fmt"

func main() {
    fmt.Println(100) // Output: 100
}
```

---

### 🔹 Floating Point  
Tipe **floating point** digunakan untuk menyimpan angka **desimal**.  

| Jenis       | Presisi | Nilai Minimum | Nilai Maksimum |
|-------------|---------|--------------|--------------|
| `float32`   | ~6-9 digit | ±1.18 × 10⁻³⁸ | ±3.4 × 10³⁸ |
| `float64`   | ~15-17 digit | ±2.23 × 10⁻³⁰⁸ | ±1.8 × 10³⁰⁸ |
| `complex64` | Real dan imaginary (float32) | - | - |
| `complex128`| Real dan imaginary (float64) | - | - |

**Contoh:**  
```go
package main

import "fmt"

func main() {
    fmt.Println(3.14) // Output: 3.14
}
```

---

### 🔹 Alias Tipe Data  
Golang juga menyediakan alias untuk beberapa tipe data numerik:  

| Alias | Tipe Asli |
|-------|----------|
| `byte` | `uint8` |
| `rune` | `int32` |
| `int`  | Minimal `int32` |
| `uint` | Minimal `uint32` |

**Contoh penggunaan alias:**  
```go
package main

import "fmt"

func main() {
    fmt.Println('A') // Output: 65 (kode ASCII untuk 'A')
}
```

---

## Tipe Data Boolean  
Tipe data **Boolean** hanya memiliki **dua nilai**, yaitu:  
- `true` → bernilai **benar**  
- `false` → bernilai **salah**  

Di Go, tipe data ini direpresentasikan dengan kata kunci `bool`.  

**Contoh:**  
```go
package main

import "fmt"

func main() {
    fmt.Println(true)  // Output: true
    fmt.Println(false) // Output: false
}
```

---

## 🔥 Kesimpulan  
- **String**: Digunakan untuk menyimpan teks.  
- **Number**: Terdiri dari Integer (`int`, `uint`) dan Floating Point (`float32`, `float64`).  
- **Boolean**: Digunakan untuk menyimpan nilai **true** atau **false**.  
