package main

import "fmt"

// Perulangan
func main() {
	dataUntukDiulang := []string{"satu", "dua", "tiga"}

	// Cara ke 1
	for index, isiDariData := range dataUntukDiulang {
		fmt.Println(index, isiDariData)
	}

	// Cara ke 2
	for index := 0; index < len(dataUntukDiulang); index++ {
		fmt.Println(index, dataUntukDiulang[index])
	}

	// Hanya Menggunakan salah satu value dari slice
	for _, isiDariData := range dataUntukDiulang {
		fmt.Println(isiDariData)
	}

	for index, _ := range dataUntukDiulang {
		fmt.Println(index)
	}

	for index := range dataUntukDiulang {
		fmt.Println(index)
	}

	// // Perulangan tak terhingga. uncomment untuk menjalankan
	// for {
	// 	fmt.Println("Tak Terhingga")
	// }
}
