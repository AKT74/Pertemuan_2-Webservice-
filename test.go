package main

import "fmt"

func main() {
    var var1, var2 int
    var operator string

    // Memasukkan variabel 1
    fmt.Print("Masukkan variabel 1: ")
    fmt.Scan(&var1)

    // Memasukkan variabel 2
    fmt.Print("Masukkan variabel 2: ")
    fmt.Scan(&var2)

    // Memilih operasi
    fmt.Print("Pilih operasi (+ untuk penjumlahan, - untuk pengurangan, * untuk perkalian, / untuk pembagian): ")
    fmt.Scan(&operator)

    // Memanggil fungsi performOperation
    result := performOperation(var1, var2, operator)

    // Menampilkan hasil
    fmt.Println("Hasil operasi:", result)
}

// Fungsi performOperation untuk melakukan operasi matematika
func performOperation(a, b int, operator string) int {
    switch operator {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        if b != 0 {
            return a / b
        } else {
            fmt.Println("Error: Pembagian dengan nol tidak diizinkan.")
            return 0
        }
    default:
        fmt.Println("Error: Operasi tidak valid.")
        return 0
    }
}
