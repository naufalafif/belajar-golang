package main

import "fmt"

func main() {
	nama := "Budi"
	ubahNama(nama)
	fmt.Println("Tanpa Pointer", nama) // Nama Tidak Berubah

	ubahNamaDenganPointer(&nama)
	fmt.Println("Dengan Pointer", nama) // Nama Berubah, karena string termasuk value type, bukan reference type. value type membutuhkan pointer untuk mengubahh data.
	// silahkan lihat gambar, variable-type.png

	dataSlice := []string{"aku", "saya", "dia", "mekera"}

	ubahSlice(dataSlice) // data berubah karena, slice adalah reference type, jadi tidak memerlukan pointer
	fmt.Println("tanpa pointer", dataSlice)

}

func ubahSlice(dataSlice []string) {
	dataSlice[0] = "telah diubah"
}

func ubahNama(nama string) {
	nama = "adi"
}

func ubahNamaDenganPointer(nama *string) {
	*nama = "adi"
}
