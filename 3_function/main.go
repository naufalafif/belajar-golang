package main

import "fmt"

// Function Biasa
func kaliDua(angka int) int {
	return angka * 2
}

func main() {
	fmt.Println("Fungsi Biasa", kaliDua(2))

	data := "Tampikan Di Fungsi Literan"
	// Function Literal / Lambda Functin / Anonymous Functin
	func(dataArgumen string) {
		fmt.Println("Fungsi Literal", dataArgumen)
	}(data)
}
